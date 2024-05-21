# Python gRPC Client

Same thing, but in Python.

```
python3 -m venv venv
source venv/bin/activate
pip3 install protobuf grpcio grpcio-tools

# make the gRPC stuff
cd .. && make helloworld-py

# start the server

# run it!
cd client-py
python3 client.py
```
