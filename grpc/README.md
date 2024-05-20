# GRPC

From [How To: Build a gRPC Server In Go](https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05)

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
export PATH="$PATH:$HOME/.local/bin:$(go env GOPATH)/bin"
```
