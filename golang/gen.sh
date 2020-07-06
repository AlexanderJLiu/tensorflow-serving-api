#! /bin/bash

PROJECT=/path/to/tensorflow-serving-api
version=2.2.0

cd ${PROJECT}/tensorflow
git checkout tags/v${version}

cd ${PROJECT}/serving
git checkout tags/${version}

cd ${PROJECT}

protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang serving/tensorflow_serving/*/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang serving/tensorflow_serving/sources/storage_path/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/framework/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/example/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/protobuf/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/stream_executor/*.proto

cd ${PROJECT}/tensorflow
git checkout master

cd ${PROJECT}/serving
git checkout master