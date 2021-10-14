gRPC example server and clients
===============================

This project sets up an example server (`/server`) written in Golang. There are two clients, one written in Golang (`/clients/go`) and the other written in Node.js (`/clients/node`) in order to demonstrate how easy it is to consume the API using gRPC.

This example server sets up a web server on localhost:8000 which implements a `GetGame` function. Users can ask for a game by title (i.e. "Uncharted 4") and then receive the first game wich matches the title - all done using gRPC.

# Requirements
- Golang
- `protoc` (on mac installable with `brew install protobuf`)

# Set up
- `go get ./...` to get the dependencies
- `go run server/server.go` to launch the web server
- To use a client accessing the API:
  - Golang:
    - `go run clients/go/client.go`
  - Node.js:
    - `cd clients/node`
    - `npm install`
    - `node index.js`
