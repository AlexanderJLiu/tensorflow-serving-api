import grpc
import tensorflow as tf
from absl import app, flags
import numpy as np

from tensorflow.core.framework import types_pb2
from tensorflow_serving.apis import predict_pb2
from tensorflow_serving.apis import prediction_service_pb2_grpc

flags.DEFINE_string(
    'server',
    '127.0.0.1:8500',
    'PredictionService host:port',
)
FLAGS = flags.FLAGS


def main(_):
    channel = grpc.insecure_channel(FLAGS.server)
    stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)

    request = predict_pb2.PredictRequest()
    request.model_spec.name = 'first_model'
    request.model_spec.signature_name = 'serving_default'
    # request.model_spec.version_label = "stable"
    request.model_spec.version.value = 0
    data = np.ones((2, 31))
    request.inputs['input_1'].CopyFrom(
        tf.make_tensor_proto(data, dtype=types_pb2.DT_INT64))

    request.output_filter.append('output_1')
    result = stub.Predict(request, 10.0)  # 10 secs timeout

    print(result)
    print(result.outputs["output_1"].float_val)


if __name__ == '__main__':
    app.run(main)