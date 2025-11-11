# Go gRPC RouteGuide Example

This project implements the official gRPC Go basics tutorial example from https://grpc.io/docs/languages/go/basics

## Overview

The RouteGuide service demonstrates all four types of gRPC service methods:

1. **Simple RPC** (`GetFeature`) - Client sends a single request, server sends a single response
2. **Server-side streaming RPC** (`ListFeatures`) - Client sends a request, server sends back a stream of responses
3. **Client-side streaming RPC** (`RecordRoute`) - Client sends a stream of requests, server sends back a single response
4. **Bidirectional streaming RPC** (`RouteChat`) - Both client and server send streams of messages to each other

## Project Structure

```
.
├── proto/
│   └── route_guide.proto          # Protocol buffer definition
├── server/
│   └── main.go                    # gRPC server implementation
├── client/
│   └── main.go                    # gRPC client implementation
├── scripts/
│   ├── generate_proto.ps1         # Generate Go code from proto
│   ├── run_server.ps1             # Run the server
│   └── run_client.ps1             # Run the client
├── go.mod                         # Go module file
└── README.md                      # This file
```

## Prerequisites

Before running this example, you need:

1. **Go** (1.20 or later) - [Download Go](https://golang.org/dl/)
2. **Protocol Buffers Compiler (protoc)** - [Download protoc](https://github.com/protocolbuffers/protobuf/releases)
   - Make sure `protoc` is in your PATH
3. **Go plugins for protoc**:

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure `$GOPATH/bin` (usually `%USERPROFILE%\go\bin` on Windows) is in your PATH.

## Setup and Run

### 1. Generate Go code from proto

From the project root directory:

```powershell
.\scripts\generate_proto.ps1
```

Or run protoc directly:

```powershell
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/route_guide.proto
```

This generates two files in the `proto/` directory:
- `route_guide.pb.go` - Protocol buffer message definitions
- `route_guide_grpc.pb.go` - gRPC service definitions

### 2. Download dependencies

```powershell
go mod tidy
```

### 3. Run the server

In one PowerShell window:

```powershell
.\scripts\run_server.ps1
```

Or:

```powershell
go run ./server
```

You should see: `gRPC server listening on [::]:50051`

### 4. Run the client

In another PowerShell window:

```powershell
.\scripts\run_client.ps1
```

Or:

```powershell
go run ./client
```

## Expected Output

The client will:

1. Call `GetFeature` for a specific point (simple RPC)
2. Call `GetFeature` for a missing feature (returns unnamed feature)
3. Call `ListFeatures` to get all features in a rectangle (server streaming)
4. Call `RecordRoute` to send random points and get route summary (client streaming)
5. Call `RouteChat` to exchange messages at various locations (bidirectional streaming)

Example output:

```
Getting feature for point (409146138, -746188906)
name:"Patriots Path, Mendham, NJ 07945, USA" location:<latitude:409146138 longitude:-746188906 >
Getting feature for point (0, 0)
location:<>
Looking for features within ...
Feature: name: "Patriots Path, Mendham, NJ 07945, USA", point:(407838351, -746143763)
...
Traversing 56 points.
Route summary: point_count:56 feature_count:2 distance:12345 elapsed_time:0
Got message First message at point(0, 1)
...
```

## Service Methods

### GetFeature (Simple RPC)
Returns information about a feature at a given position.

### ListFeatures (Server-side streaming)
Returns all features within a given bounding rectangle. Results are streamed.

### RecordRoute (Client-side streaming)
Accepts a stream of Points on a route being traversed, returns a RouteSummary when complete.

### RouteChat (Bidirectional streaming)
Accepts a stream of RouteNotes sent while traversing a route, while receiving other RouteNotes (e.g., from other users).

## Notes

- The module path is set to `github.com/matsb/grpc` in `go.mod` and the proto file's `go_package` option
- If you want to change the module path, update both files and regenerate the proto code
- The server includes hardcoded sample features for demonstration
- Port 50051 is used by default (standard gRPC examples port)

## Learn More

- [gRPC Go Basics Tutorial](https://grpc.io/docs/languages/go/basics/)
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers Go Tutorial](https://protobuf.dev/getting-started/gotutorial/)
