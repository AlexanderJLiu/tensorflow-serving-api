<div align="center">
    <img src="https://raw.githubusercontent.com/AlexanderJLiu/tensorflow-serving-api/master/images/python.png"
        width="200" height="200">
</div>

# TensorFlow Serving API for Python

1. Because there are pre-generated python API of tensorflow serving, We don't have to generate it ourselves.
2. Run the following command to install python package.

   ```shell
   pip install tensorflow-serving-api==2.2.0
   ```

3. Run the code in main.py for sending a grpc request to tensorflow serving.

   ```shell
   python main.py
   ```

4. Note that the model metadata is the same as the model in golang [README.md](https://github.com/AlexanderJLiu/tensorflow-serving-api/tree/master/golang#make-a-grpc-request).
