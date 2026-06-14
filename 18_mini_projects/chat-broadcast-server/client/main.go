// Command client is a minimal TCP chat client for chat-broadcast-server.
//
// It dials the server, then runs two concurrent copy loops:
//   - socket -> stdout: print whatever the server broadcasts
//   - stdin  -> socket: send whatever you type
//
// Try it (server must be running on :8080):
//
//	go run ./18_mini_projects/chat-broadcast-server/server   # terminal 1: server
//	go run ./18_mini_projects/chat-broadcast-server/client   # terminal 2+: clients
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// -addr lets you point at a different host/port; defaults to the server's.
	addr := flag.String("addr", "localhost:8080", "chat server address host:port")
	flag.Parse()

	// Establish the TCP connection.
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("dial %s: %v", *addr, err)
	}
	defer conn.Close()
	log.Printf("connected to %s — type a message and press Enter (Ctrl+C / Ctrl+D to quit)", *addr)

	// done signals when the server side closes (EOF on the socket), so the
	// program can exit instead of hanging on stdin forever.
	done := make(chan struct{})

	// Reader goroutine: stream everything the server sends straight to stdout.
	// io.Copy loops until the connection returns EOF or errors.
	go func() {
		// io.Copy returns the bytes copied and the first error; we only care
		// that it has finished, which means the server closed the connection.
		_, _ = io.Copy(os.Stdout, conn)
		close(done)
	}()

	// Main goroutine: copy typed input from stdin to the socket. This blocks
	// until stdin hits EOF (Ctrl+D) or the read fails.
	go func() {
		_, _ = io.Copy(conn, os.Stdin)
		// Closing the connection unblocks the reader goroutine's io.Copy.
		conn.Close()
	}()

	// Exit once the server closes the connection (or we closed it above).
	<-done
}
