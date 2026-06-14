// Command chat-broadcast-server is a TCP chat server. A single "broadcaster"
// goroutine owns the set of connected clients and fans every incoming message
// out to all of them. Because only one goroutine ever touches the client set,
// no mutex is needed — the channels serialize all access for us.
//
// Try it:
//
//	go run ./18_mini_projects/chat-broadcast-server/server
//	# in other terminals:
//	go run ./18_mini_projects/chat-broadcast-server/client   # or: nc localhost 8080
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// client is a send-only channel of messages destined for one connection.
// Using chan<- string (send-only) lets the broadcaster send to a client but
// makes it a compile error for the broadcaster to accidentally receive from it.
type client chan<- string

// broadcaster is the single goroutine that owns the `clients` map. It reacts to
// three events via a select loop: a client joining, a client leaving, and a new
// message to fan out. Serializing these in one goroutine is what removes the
// need for any locking on the shared map.
func broadcaster(incoming chan client, leaving chan client, messages chan string) {
	// The set of currently-connected clients. Only this goroutine reads or
	// writes it, so it is race-free without a mutex.
	clients := map[client]bool{}
	for {
		select {
		case c := <-incoming:
			// A new connection registered itself.
			clients[c] = true

		case c := <-leaving:
			// A connection is gone: drop it from the set and close its channel.
			// Closing ends the writer loop in handleConn (its `for range ch`).
			delete(clients, c)
			close(c)

		case msg := <-messages:
			// Fan the message out to every connected client.
			for c := range clients {
				// Non-blocking send: if a client's buffered channel is full
				// (a slow reader), we fall through to default and skip it rather
				// than letting one slow client stall the whole broadcaster.
				select {
				case c <- msg:
				default:
					// slow client — skip or disconnect
				}
			}
		}
	}
}

// handleConn serves a single TCP connection. It registers the connection with
// the broadcaster, spawns a reader that forwards inbound lines to `messages`,
// and itself acts as the writer that flushes broadcast messages back to the
// socket.
func handleConn(conn net.Conn, incoming, leaving chan client, messages chan string) {
	// Per-client buffered channel. The buffer (8) absorbs short bursts so a
	// briefly-busy writer doesn't immediately get skipped by the broadcaster's
	// non-blocking send.
	ch := make(chan string, 8)

	// Announce ourselves to the broadcaster...
	incoming <- ch
	// ...and make sure we deregister when this function returns (connection
	// closed or read error). The broadcaster will then close(ch).
	defer func() { leaving <- ch }()

	// Reader goroutine: read the socket line by line and publish each line to
	// the shared messages channel for fan-out. When the client disconnects,
	// Scan returns false and this goroutine exits.
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			messages <- scanner.Text()
		}
	}()

	// Writer loop (runs on this goroutine): drain our channel and write each
	// broadcast message to the socket. The loop ends when the broadcaster
	// closes ch, which is what triggers the deferred leaving send to complete.
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Three unbuffered channels connect the handlers to the broadcaster.
	incoming := make(chan client) // a client joined
	leaving := make(chan client)  // a client left
	messages := make(chan string) // a line to broadcast

	// Start the single broadcaster goroutine.
	go broadcaster(incoming, leaving, messages)

	// Listen for TCP connections on port 8080.
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("chat server listening on :8080")

	// Accept loop: each accepted connection is handled by its own goroutine,
	// so many clients can be served concurrently.
	for {
		conn, err := ln.Accept()
		if err != nil {
			// Transient accept error — log and keep serving.
			continue
		}
		go handleConn(conn, incoming, leaving, messages)
	}
}
