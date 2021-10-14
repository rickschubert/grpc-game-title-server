gRPC example server and clients
===============================

This project demonstrates how [gRPC](https://grpc.io/) to implement APIs on web servers. Unlike in REST APIs, all of the exposed web server methods are defined in a contract [which the .proto file forms](gametitles/game_titles.proto).

We have an example server (`/server`) written in Golang as well as two clients, one in Golang (`/clients/go`) and one in Node.js (`/clients/node`). This demonstrates how easy it is to consume the API using gRPC and how the gRPC framework auto-generates consumer libraries and provides them to us for free.

The example server starts its API on localhost:8000 and implements a `GetGame` function. Consumers can request this API for a game by title (i.e. "Uncharted 4") and then receive the first game wich matches the title - all done using gRPC.

# Requirements
- Golang
- `protoc` (on mac installable with `brew install protobuf`)

# Set up
- `go get ./...` to get the dependencies
- `go run server/server.go` to launch the web server
- To use a client accessing the API exposed by the web server:
  - Golang:
    - `go run clients/go/client.go`
  - Node.js:
    - `cd clients/node`
    - `npm install`
    - `node index.js`

# Development
The protobuf Go stubs are created using the script `bash compileproto.sh`. Whenever the `.proto` file is changed, this task needs to be re-run.
