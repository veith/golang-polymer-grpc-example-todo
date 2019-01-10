#! /bin/bash

if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    echo "./scripts/build.sh protoname"
    exit 1
fi

if [ ! -d $1 ]
then
     mkdir $1
fi

echo "Generating $1.pb.go"
echo "Generating $1.pb.gw.go"
echo "Generating doc/$1.swagger.json"

protoc -I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/googleapis/googleapis \
--gogofast_out=\
Mgoogle/type/date.proto=google.golang.org/genproto/googleapis/type/date,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:./$1/ \
--swagger_out=logtostderr=true:./doc \
--grpc-gateway_out=logtostderr=true:./$1/ \
./$1.proto