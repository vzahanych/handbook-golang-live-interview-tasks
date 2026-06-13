# Time-zone day boundary

## Live interview task
Given an instant and an IANA zone, return the start of that local calendar day as an instant.

## Candidate solution

```go
loc, err := time.LoadLocation(zone)
if err != nil { return time.Time{}, err }
local := instant.In(loc)
return time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, loc), nil
```

## Interview notes / pitfalls
- A local day is not always 24 hours around daylight-saving transitions.
- Store instants and zone identifiers separately when future local scheduling matters.
