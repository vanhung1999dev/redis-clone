# GoRedis

A simple Redis-compatible in-memory key-value store implemented in Go.  
Supports basic Redis commands like `SET`, `GET`, `EXISTS`, `DEL`, `CLIENT`, and `HELLO`.

## Features

- In-memory key-value storage
- Basic Redis protocol support (RESP)
- Expiration with `SET EX`
- Command options: `NX`, `XX`
- Tested with `redis-cli`

---

## Commands Supported

- `SET key value [EX seconds] [NX|XX]`
- `GET key`
- `EXISTS key`
- `DEL key`
- `CLIENT <value>`
- `HELLO <value>`

---

## Getting Started

### Prerequisites

- Go 1.18+
- [redis-cli](https://redis.io/docs/getting-started/installation/) installed (for testing)

---

### Build & Run

```bash
# Clone the repo
git clone https://github.com/vanhung1999dev/redis-clone
cd redis-clone

# Run the server

```

> By default, the server listens on `localhost:5001`. You can change it with the `--listenAddr` flag:

```bash
go run . --listenAddr :6379
```
> Start by using Makefile. You should install it in your computer to can run project.
```
make run
```

---

### Example with `redis-cli`

Start your server:

```bash
go run .
```

In another terminal, connect using `redis-cli`:

```bash
redis-cli -p 5001
```

#### Try some commands:

```bash
127.0.0.1:5001> SET foo bar
OK

127.0.0.1:5001> GET foo
"bar"

127.0.0.1:5001> EXISTS foo
(integer) 1

127.0.0.1:5001> DEL foo
(integer) 1

127.0.0.1:5001> GET foo
(nil)
```

#### Set with expiration:

```bash
127.0.0.1:5001> SET temp "will-expire" EX 5
OK

127.0.0.1:5001> GET temp
"will-expire"

# Wait 5 seconds...

127.0.0.1:5001> GET temp
(nil)
```

---

## Project Structure

```
.
├── main.go          # Entry point, server setup
├── peer.go          # TCP server, peer connection management
├── proto.go         # RESP command parsing
├── types.go         # Command definitions and helpers
├── keyval.go        # In-memory key-value store
└── README.md
```

---

## How It Works

### Protocol

This server uses the Redis Serialization Protocol (RESP), the same used by official Redis. That's why you can interact with it using `redis-cli`.

### Memory Store

It uses a simple in-memory map protected by mutexes for thread safety. TTL (time to live) expiration is supported by checking timestamps on `GET`.

---

## TODOs / Possible Enhancements

- [ ] Add support for `PING`, `INCR`, `MGET`, etc.
- [ ] Add persistence (RDB or AOF)
- [ ] Background TTL cleanup
- [ ] Pub/Sub support
- [ ] Authentication with `AUTH`
- [ ] Benchmarks and load testing

---

## License

MIT License.  
Feel free to use, modify, and share.

---

## Author

**Your Name**  
[GitHub](https://github.com/vanhung1999dev)
