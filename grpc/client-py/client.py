import grpc
import logging

# this is ugly ;-)
import os
import sys
parent_dir_name = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
sys.path.append(parent_dir_name + "/helloworld")
import helloworld_pb2
import helloworld_pb2_grpc

from concurrent import futures
import time

def run():
    # Create a gRPC channel to the server
    with grpc.insecure_channel('localhost:50051') as channel:
        # Create a stub (client)
        stub = helloworld_pb2_grpc.HelloWorldServiceStub(channel)
        
        # Create a request
        request = helloworld_pb2.HelloWorldRequest()
        
        # Make the call
        try:
            response = stub.SayHello(request, timeout=1)  # 1 second timeout
            print("Response from gRPC server's SayHello function: ", response.message)
        except grpc.RpcError as e:
            logging.error(f"Error calling SayHello: {e}")
            
if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    run()
