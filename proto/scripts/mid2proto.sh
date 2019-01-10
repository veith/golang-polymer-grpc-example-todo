#! /bin/bash


echo "Generating $1.proto"
simple-generator -d $1.mid -t scripts/proto.tmpl > $1.proto
