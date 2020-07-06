<div align="center">
    <img src="https://raw.githubusercontent.com/AlexanderJLiu/tensorflow-serving-api/master/images/golang.png"
        width="150" height="200">
</div>

# Tensorflow Serving API for Golang

## Prerequisites

### golang

1. Download the golang archive [here](https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz).

   ```shell
   wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
   ```

2. Extract it into /usr/local, creating a go tree in /usr/local/go.

   ```shell
   tar zxvf go1.14.4.linux-amd64.tar.gz -C /usr/local
   ```

3. Add `/usr/local/go/bin` to the PATH environment variable. You can do this by adding this line to your `/etc/profile`.

   ```shell
   export PATH=$PATH:/usr/local/go/bin
   ```

4. Test your installation

   ```shell
   go version
   ```

### protoc

1. Protoc is used to compile protocol buffer files.
2. Go to the release page of [protobuf](https://github.com/protocolbuffers/protobuf/releases) and download the latest `protoc` binary tarball.

   ```shell
   wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip
   ```

3. Extract it into /usr/local, creating a protoc tree in /usr/local/protoc.

   ```shell
   unzip protoc-3.12.3-linux-x86_64.zip -d /usr/local/protoc
   ```

4. Add `/usr/local/protoc/bin` to the PATH environment variable. You can do this by adding this line to your `/etc/profile`.

   ```shell
   export PATH=$PATH:/usr/local/protoc/bin
   ```

5. Test your installation.

    ```shell
    protoc --version
    ```

### protoc-gen-go

1. protoc-gen-go is a plugin of protoc to generate `go` files.
2. Run the following command to install the go protocol buffers plugin.

   ```shell
   go get -u google.golang.org/protobuf/cmd/protoc-gen-go
   # or
   go get -u github.com/golang/protobuf/protoc-gen-go
   ```

3. The compiler plugin protoc-gen-go will be installed in `$GOBIN`, defaulting to `$GOPATH/bin`. It must be in your `$PATH` for the protocol compiler protoc to find it.

### grpc(optional)

1. Grpc plugin is used to wrap the request uri for ease of use. You could also write the original code for sending request to tensorflow serving without installing this pulgin.
2. I suggest you better install it. Run the following command for installation.

   ```shell
   go get -u google.golang.org/grpc
   ```

## Generate Go API File

### Swith to corresponding branch/tag

1. Clone this project to your local dir with the following command.

   ```shell
   git clone https://github.com/AlexanderJLiu/tensorflow-serving-api.git
   git submodule init
   git submodule update
   # or
   git clone --recurse-submodules https://github.com/AlexanderJLiu/tensorflow-serving-api.git
   ```

2. You should swith the source code of `tensorflow` and `serving` to the corresponding branch/tag when gengerating the go files.
3. Take version 2.2.0 as an example. Run the following commands to switch branch/tag.

   ```shell
   cd /path/to/tensorflow-serving-api/tensorflow
   git checkout tags/v2.2.0
   cd /path/to/tensorflow-serving-api/serving
   git checkout tags/2.2.0
   ```

### Generate go files

Run the following commands to generate go files to specific dir `golang`.

```shell
cd /path/to/tensorflow-serving-api
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang serving/tensorflow_serving/*/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang serving/tensorflow_serving/sources/storage_path/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/framework/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/example/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/core/protobuf/*.proto
protoc -I=serving -I=tensorflow --go_out=plugins=grpc:golang tensorflow/tensorflow/stream_executor/*.proto
```

In above commands, `-I` specify the directory in which to search for `proto imports`. And it can be specified multiple times. `--go_out` specify the directory in which to save go files and also specify the plugins used for generation.

After executing the above commands, there will be two directory created in `golang` dir, one is `github.com` and another is `tensorflow_serving`. That's because all proto files of tensorflow have the `option go_package` in them, which specify the output dir of go files. However proto files of serving don't have the `go_package` option, it just create same output dir as the source code.

Also you can use the `gen.sh` script to generate go files.

## Make a grpc request

1. Initialize three go projects in `golang` folder with go module.

    ```shell
    cd github.com/tensorflow/tensorflow
    go mod init github.com/tensorflow/tensorflow
    cd tensorflow_serving
    go mod init tensorflow_serving
    cd golang
    go mod init github.com/alex/tensorflow-serving-api-go
    ```

2. Modify the `go.mod` file in `golang` dir with the following statement.

    ```go
    replace (
        github.com/tensorflow/tensorflow => ./github.com/tensorflow/tensorflow
        tensorflow_serving => ./tensorflow_serving
    )
    ```

3. If you encouter following `cycle import` error, you should delete the `logging.pb.go` file in `tensorflow_serving/core` dir and `prediction_log.pb.go` file in `tensorflow_serving/apis` dir.

   ```shell
    import cycle not allowed
    package github.com/alex/tensorflow-serving-api-go
        imports tensorflow_serving/apis
        imports tensorflow_serving/core
        imports tensorflow_serving/apis
    ```

4. Suppose I have a model named `first_model` served in tensorflow serving. Model metadata is shown as follows:

    ```shell
    $ curl http://localhost:8501/v1/models/first_model/versions/0/metadata
    {
        "model_spec": {
            "name": "first_model",
            "signature_name": "",
            "version": "0"
        },
        "metadata": {
            "signature_def": {
                "signature_def": {
                    "serving_default": {
                        "inputs": {
                            "input_1": {
                                "dtype": "DT_INT64",
                                "tensor_shape": {
                                    "dim": [
                                        {
                                            "size": "-1",
                                            "name": ""
                                        },
                                        {
                                            "size": "31",
                                            "name": ""
                                        }
                                    ],
                                    "unknown_rank": false
                                },
                                "name": "serving_default_input_1:0"
                            }
                        },
                        "outputs": {
                            "output_1": {
                                "dtype": "DT_FLOAT",
                                "tensor_shape": {
                                    "dim": [
                                        {
                                            "size": "-1",
                                            "name": ""
                                        },
                                        {
                                            "size": "1",
                                            "name": ""
                                        }
                                    ],
                                    "unknown_rank": false
                                },
                                "name": "StatefulPartitionedCall:0"
                            }
                        },
                        "method_name": "tensorflow/serving/predict"
                    },
                    "__saved_model_init_op": {
                        "inputs": {},
                        "outputs": {
                            "__saved_model_init_op": {
                                "dtype": "DT_INVALID",
                                "tensor_shape": {
                                    "dim": [],
                                    "unknown_rank": true
                                },
                                "name": "NoOp"
                            }
                        },
                        "method_name": ""
                    }
                }
            }
        }
    }
    ```

5. Then I use the code in `main.go` to send a grpc request to the model above. And the response is shown as follows:

    ```protobuf
    model_spec:{name:"first_model" version:{} signature_name:"serving_default"} outputs:{key:"output_1" value:{dtype:DT_FLOAT tensor_shape:{dim:{size:2} dim:{size:1}} float_val:0.77852035 float_val:0.77852035}}
    ```

## References

1. [go installation instructions](https://golang.org/doc/install)
2. [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
