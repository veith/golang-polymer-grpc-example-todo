# golang-polymer-grpc-example-todo
Is a training project for my students at [Web Professionals.ch](https://web-professionals.ch/) to train how to build a 
event-driven RESTFul microservice architecture with **golang**, **grpc**, **nats**, **aws**, **google cloud engine**, **web-components**,... 

## Prereqs

```
brew install protobuf

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u moul.io/protoc-gen-gotemplate

go get github.com/gogo/protobuf/protoc-gen-gogofast
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto

go get github.com/oklog/ulid
go get upper.io/db.v3
go get upper.io/db.v3/sqlite

```

OR to build protobuf from source, the following tools are needed:

- autoconf
- automake
- libtool
- make
- g++
- unzip

https://github.com/protocolbuffers/protobuf/blob/master/src/README.md


Install protobuf
```
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

## built in Generators
To use `scripts/regenerateProto.sh for generating the proto stubs and the REST-grpc-REST transcoder from proto files install
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/micro/protobuf/protoc-gen-go

go get github.com/gogo/protobuf/protoc-gen-gogofast 
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto

go get github.com/oklog/ulid
go get upper.io/db.v3
go get upper.io/db.v3/sqlite

```

To use `scripts/midDemo.sh` to generate protos install 
```
go get github.com/veith/simple-generator
```