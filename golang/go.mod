module github.com/alex/tensorflow-serving-api-go

go 1.14

replace (
	github.com/tensorflow/tensorflow => ./github.com/tensorflow/tensorflow
	tensorflow_serving => ./tensorflow_serving
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/tensorflow/tensorflow v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0
	tensorflow_serving v0.0.0-00010101000000-000000000000
)
