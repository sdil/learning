# Introduction
Go is a really awesome programming language. My daily work does not involve in writing Go so I am learning Go on my own

# What I have learned
- [todolist-mongo-go](todolist-mongo-go)
  - Basic Go program structures, including data structures in Go. Learned how to use maps & slices.
  - How to connect Go to mongo (using mgo driver)
  - How to write basic Go HTTP server using Gorilla Mux
  - How to add CORS headers in Gorilla Mux
- [todolist-mysql-go](todolist-mysql-go)
  - How to use ORM in Golang with GORM
  - Wrote a [Medium article](https://medium.com/better-programming/build-a-simple-todolist-app-in-golang-82297ec25c7d) on this
- [URL Shortener](url-shortener)
  - Hexagonal / Port and adapter architecture in Golang
  - Work with bigger codebases and multiple files
  - Use Echo web server
  - Use `context.Background()` to do timeout
  - Initialize `go mod`
  - Encountered `cyclic import` issues

# What I wish to learn
- Go unittest & integration testing
- Goroutine for concurrency applicatoin
- ORM connection with relational DBs (MySQL & PostgreSQL) **DONE**
- Write commandline apps
- Write a bigger codebase **DONE**
- Write a server that serve different protocols:
  - gRPC
  - Websocket
  - GraphQL
- Understand established OSS codes written in Go (eg. BoltDB, Dgraph, InfluxDB, etc.) to understand how people are writing Go in production
