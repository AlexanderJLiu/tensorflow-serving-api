package main

import (
	"context"
	"log"
	apis "tensorflow_serving/apis"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	"google.golang.org/grpc"
)

var (
	// TensorFlow Serving grpc address.
	address = "127.0.0.1:8500"
)

func main() {
	// Create a grpc request.
	request := &apis.PredictRequest{
		ModelSpec: &apis.ModelSpec{},
		Inputs:    make(map[string]*tensor_go_proto.TensorProto),
	}
	request.ModelSpec.Name = "first_model"
	request.ModelSpec.SignatureName = "serving_default"
	// request.ModelSpec.VersionChoice = &apis.ModelSpec_VersionLabel{VersionLabel: "stable"}
	request.ModelSpec.VersionChoice = &apis.ModelSpec_Version{Version: &wrappers.Int64Value{Value: 0}}
	request.Inputs["input_1"] = &tensor_go_proto.TensorProto{
		Dtype: types_go_proto.DataType_DT_INT64,
		TensorShape: &tensor_shape_go_proto.TensorShapeProto{
			Dim: []*tensor_shape_go_proto.TensorShapeProto_Dim{
				&tensor_shape_go_proto.TensorShapeProto_Dim{
					Size: int64(2),
				},
				&tensor_shape_go_proto.TensorShapeProto_Dim{
					Size: int64(31),
				},
			},
		},
		Int64Val: []int64{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		},
	}

	// Create a grpc connection.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("couldn't connect: %s", err.Error())
	}
	defer conn.Close()

	// Wrap the grpc uri with client.
	client := apis.NewPredictionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Send the grpc request.
	response, err := client.Predict(ctx, request)
	if err != nil {
		log.Fatalf("couldn't get response: %v", err)
	}
	log.Printf("%+v", response)
}
