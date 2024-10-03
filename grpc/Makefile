.PHONY: generate

helloworld:
	cd helloworld && \
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
    	helloworld.proto

# not sure if this is right for Python -- protoc is complaining
# about missing plugins
helloworld-py:
	cd helloworld && \
	python3 -m grpc_tools.protoc \
		--proto_path=. \
		--python_out=. \
		--grpc_python_out=. \
		helloworld.proto
