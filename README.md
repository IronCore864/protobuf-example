# protobuf-example

Golang protobuf example

## Protobuf Compiler Installation

The simplest way to install the protocol compiler is to download a pre-built binary from the official release page:

https://github.com/protocolbuffers/protobuf/releases

Make sure the protoc binary is in your path.

## Install the Proto Package

The simplest way is to run:

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Compile

```
SRC_DIR=$(pwd)/pb
DST_DIR=$(pwd)/pb
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

## Build Project

```
go build

```

## Write/Read Example

```
./protobuf-example write
./protobuf-example read
```
