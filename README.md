# Getting started with [Protobuffs](https://developers.google.com/protocol-buffers/)

Protobuffs is a convinient way to define you public/private API and auto
generate code handlers from this definition. This is a super simple **Go**
blueprint which you can use to start with protobuffs. It consists of server
and client. It shows a minimal setup of protobuffs and the flow of generating
models from protobuffs definitions.


## Usage
Get protobuff compiler from [here](https://github.com/google/protobuf/releases)

```
# Get protobuff go extension
go get -u github.com/golang/protobuf/protoc-gen-go

# Get the repo
go get -u github.com/itsankoff/protobuffs/

# Make sure GOPATH/bin is in you PATH, because protobuff compiler needs to
# find protoc-gen-go in order to compile the definitions to go models.

# Go to $GOPATH/src/github.com/itsankoff/protobuff/
# Executing the following command will generate go models from the definitions
# within $(pwd)/def/definitions folder.
protoc --proto_path=$(pwd)/def/definitions --go_out=$(pwd)/def $(pwd)/def/definitions/*.proto

# Start server and client
go run ./cmd/server/main.go
go run ./cmd/client/main.go
```
