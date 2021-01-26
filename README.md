### GRPC_TEST

### this is a test project for grpc

1.  fist download package

```
go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
```
 
2. generate `.proto` file

```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

ex.
protoc --go_out=plugins=grpc:. pb/*.proto
 ```