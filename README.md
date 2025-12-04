# 🔧 Go Projects for Learning and Practice

This repository contains a collection of small Go projects designed to help me learn and master various aspects of the Go programming language. Each project focuses on a different area, from concurrency and networking to file handling and API development.

---

## 📡 1. Chat Room Server

**Goal**: Build a real-time chat server using WebSockets.

### Key Concepts:
- Concurrency with Goroutines
- Message broadcasting with Channels
- WebSocket handling ('net')

### Bonus Features:
- User nicknames
- Chat history log

---

## 📝 2. CLI Task Manager

**Goal**: A command-line tool to manage personal tasks with JSON-based persistence.

### Key Concepts:
- Structs and Slices
- File I/O (`os`, `ioutil`)
- JSON Encoding/Decoding
- Command-line flags or `cobra` CLI framework

### Bonus Features:
- Task filtering and search
- Tags and due dates
- Color-coded output

---

## 🗒️ 3. REST API for a Notes App

**Goal**: A RESTful API for creating and managing notes.

### Key Concepts:
- HTTP server using `Gin` or `gorilla/mux`
- CRUD operations with JSON
- Request/response validation
- Middleware and graceful shutdown

### Bonus Features:
- SQLite or BoltDB integration
- Token-based authentication (JWT)
- API documentation (Swagger)

---

## 🖼️ 4. Image Resizer Microservice

**Goal**: A simple HTTP service that resizes uploaded images.

### Key Concepts:
- Image processing with `image` and `github.com/nfnt/resize`
- File uploads (`multipart/form-data`)
- HTTP handling and routing
- Goroutines and buffered channels for concurrency control

### Bonus Features:
- Support for multiple image formats (.jpg, .png, .webp)
- Image caching
- Adjustable resize parameters via query strings

---

## 📄 5. Markdown to HTML Converter

**Goal**: Convert Markdown files into HTML pages.

### Key Concepts:
- File reading and writing
- Parsing Markdown (`github.com/russross/blackfriday`)
- HTML templating with `html/template`
- Directory traversal with `filepath.Walk`

### Bonus Features:
- Watcher for automatic rebuilds on file change
- Custom themes/templates
- Table of contents generation

---

## 📄 6. Minimal Redis Clone

**Goal**: Build a tiny in-memory key–value store inspired by Redis to learn about networking, concurrency, and commandparsing.

### Key Concepts:
- TCP servers using Go’s net package
- Handling multiple clients with goroutines
- Parsing simple RESP (Redis Serialization Protocol)
- Implementing core commands (GET, SET, DEL, PING)
- Safe concurrent access with sync.RWMutex or sync.Map


---

## 🧠 Purpose

Each project in this repository is intended to help me:
- Learn core Go concepts
- Practice real-world coding patterns
- Explore Go libraries and tools
- Build a solid foundation for future Go development


---
