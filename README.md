# go-systems-from-scratch

> Building systems from the ground up in Go — no frameworks, no tutorials, no copy-paste.  
> Every line written and understood before it gets committed.

**Author:** Hashim — [@KillerPand34973](https://x.com/KillerPand34973)  
**Started:** May 2026 | Hyderabad, India

---

## What this is

Six projects that build on each other. Each one teaches a core CS concept by forcing you to implement it — not read about it.

By the end: a working Redis-like server built from a raw TCP socket up. Tested with actual `redis-cli`.

No Gin. No frameworks. No ORMs. Just Go's standard library and the things I build myself.

---

## Projects

### ✅ / 🔲 Status

| # | Project | Status | CS Concept Unlocked |
|---|---------|--------|----------------------|
| 0 | `memview` | 🔲 | Computer Architecture — CPU alignment, memory padding |
| 1 | `tinytcp` | 🔲 | Computer Networks — raw TCP sockets, concurrent connections |
| 2 | `httpfromscratch` | 🔲 | Networks + Compiler — HTTP/1.1 as a text protocol, parsing |
| 3 | `mini-kvstore` | 🔲 | DBMS — Write-Ahead Log, crash recovery, race conditions |
| 4 | `goroutine-pool` | 🔲 | OS — semaphores, bounded concurrency, worker pools |
| 5 | `mini-redis` | 🔲 | Everything combined — RESP protocol, persistence, TTL |

---

## Project 0 — memview

**The puzzle:** Why does reordering struct fields change the struct's total size?

```go
type BadLayout struct {
    A bool    // 1 byte
    B int64   // 8 bytes
    C bool    // 1 byte
}
// unsafe.Sizeof(BadLayout{}) = 24 — not 10. Why?
```

Run it and find out.

**What it teaches:** CPU alignment padding. The compiler adds invisible bytes between fields so each type starts at the right memory address. This is why field order matters in high-performance Go code.

---

## Project 1 — tinytcp

**The puzzle:** Can two programs talk to each other with no framework at all?

Two terminals. One server. One client. Raw `net.Listen` and `net.Dial`. No HTTP. No libraries.

**What it teaches:** What a socket actually is. What `Accept()` does while it blocks. Why goroutines are spawned per connection. Every web framework you've used is this underneath.

---

## Project 2 — httpfromscratch

**The puzzle:** What does an HTTP request actually look like as bytes?

```
GET /hello HTTP/1.1\r\n
Host: localhost:8080\r\n
\r\n
```

Parse it manually. Respond manually. Open a browser. See your response.

**What it teaches:** HTTP is just a text format on top of TCP. After this, `net/http` has no magic left in it.

---

## Project 3 — mini-kvstore

**The puzzle:** Can a map inside a struct be a database?

Yes. Then: what happens when the program crashes? Add a Write-Ahead Log. Survive restarts. Find a race condition with `go run -race`. Fix it with a mutex.

**What it teaches:** WAL persistence (how every database survives crashes), race conditions, mutual exclusion. The D in ACID — durability.

---

## Project 4 — goroutine-pool

**The puzzle:** Can I run exactly 5 workers at once, no matter how many tasks?

100 tasks. Max 5 goroutines running simultaneously. A buffered channel as a semaphore. `sync.WaitGroup` to know when all 100 are done.

**What it teaches:** The semaphore pattern (Dijkstra's P/V operations, in Go). Why database connection pools exist. What killed a connection pool I once debugged at work.

---

## Project 5 — mini-redis

**The puzzle:** Does `redis-cli` work with my server?

TCP server on port 6379. RESP protocol parsed manually. SET, GET, DEL, EXISTS, KEYS. WAL persistence from Project 3. TTL/expiry. Concurrent clients from Project 4.

If `redis-cli` talks to it, the protocol is correct. That's the test.

**What it teaches:** How Redis actually works at its core. Why it's fast. What RESP is. Lazy expiry. Everything from the previous 5 projects, combined.

---

## How I'm building this

- One puzzle at a time. One session = 15 minutes.
- Code must **run** before moving to the next puzzle. Not "I understand it" — runs.
- `go run -race` on every concurrent project before calling it done.
- Every commit message describes what broke or what I learned — not "update" or "fix".

Daily progress on Twitter: [@KillerPand34973](https://x.com/KillerPand34973)

---

## Sample commit messages from this repo

```
memview — BadLayout is 24 bytes not 10, found out why (alignment padding)
tinytcp p2 — client sends, server reads, both terminals work
httpfromscratch p3 — browser shows my response, Content-Length was wrong first
mini-kvstore p5 — found race with -race flag, fixed with RWMutex
goroutine-pool — semaphore blocks correctly at 5, verified with counter
mini-redis — redis-cli SET and GET work, RESP parser correct
```

---

## Connect

- **Twitter:** [@KillerPand34973](https://x.com/KillerPand34973) — daily build log
- **LinkedIn:** [syed-hashim721](https://linkedin.com/in/syed-hashim721)
- **Learning journal:** [Golang-Reboot-2026](https://github.com/Hashim-777x/Golang-Reboot-2026)
