#! /bin/bash

PROJECT=/path/to/tensorflow-serving-api
version=2.2.0

cd ${PROJECT}/tensorflow
git checkout tags/v${version}

cd ${PROJECT}/serving
git checkout tags/${version}

cd ${PROJECT}

protoc -I=serving -I=tensorflow --plugin=/usr/local/protoc/bin/protoc-gen-grpc-java --grpc-java_out=java --java_out=java serving/tensorflow_serving/*/*.proto
protoc -I=serving -I=tensorflow --plugin=/usr/local/protoc/bin/protoc-gen-grpc-java --grpc-java_out=java --java_out=java serving/tensorflow_serving/sources/storage_path/*.proto

cd ${PROJECT}/tensorflow
git checkout master

cd ${PROJECT}/serving
git checkout master