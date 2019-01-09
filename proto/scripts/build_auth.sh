#! /bin/bash

# https://github.com/gogo/protobuf/issues/325

protoc -I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/googleapis/googleapis/google/type \
--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:./auth/ \
--swagger_out=logtostderr=true:./doc \
--grpc-gateway_out=logtostderr=true:./auth/ \
./auth.proto