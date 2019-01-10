#! /bin/bash


echo "Generating proto for $1"
simple-generator -d $1.mid -t scripts/proto.tmpl > $1.proto
