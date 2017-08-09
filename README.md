# Getting started with [Protobuffs](https://developers.google.com/protocol-buffers/)

Protobuffs is a convinient way to define you public/private API and auto
generate code handlers from this definition. This is a super simple **Go**
blueprint which you can use to start with protobuffs. It consists of server
and client. It shows a minimal setup of protobuffs and the flow of generating
models from protobuffs definitions.


## Usage
i
Get protobuff compiler from [here](https://github.com/google/protobuf/releases)

```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/itsankoff/protobuffs/

# make sure GOPATH/bin is in you PATH, because protobuff compiler needs to
# find protoc-gen-go in order to compile the definitions to go models

# go to /path/to/workspace/github.com/itsankoff/protobuff/def/definitions
# to first generate models from protobuffs definitions
protoc --go_out=../ *

# go to /path/to/workspace/github.com/itsankoff/protobuff/
go run ./cmd/server/main.go
go run ./cmd/client/main.go
```
