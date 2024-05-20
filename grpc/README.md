# gRPC

Teaching myself gRPC.

Lots of tutorials out there, some more confusing than others.  I found [How To: Build a gRPC Server In Go](https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05) to be pretty straightforward.

## TL;DR

```
make helloworld
go run server/main.go
go run client/main.go
```

## Notes to Future Self

- I think I originally installed `protoc-gen-go` and `protoc-gen-go-grpc` using `brew` which put me on the wrong path.
- This is maybe all that I needed to do for dependencies (besides the obvious like `go` and `make`):
```
brew install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Example

### Server

```
➜  grpc git:(main) ✗ go run server/main.go
2024/05/20 10:27:31 gRPC server listening at [::]:50051
```

### Client

```
➜  grpc git:(main) go run client/main.go
2024/05/20 10:51:23 Response from gRPC server's SayHello function: Hello world, the time is 2024-05-20 10:51:23
➜  grpc git:(main) go run client/main.go
2024/05/20 10:51:25 Response from gRPC server's SayHello function: Hello world, the time is 2024-05-20 10:51:25
```
