.PHONY: generate

# generate1:
# 	protoc *.proto \
# 		--go_out=. \
# 		--go_opt=paths=source_relative \
# 		--proto_path=.

# generate2:
# 	protoc api/v1/*.proto \
# 		--go-grpc_out=. \
# 		--go-grpc_opt=paths=source_relative \
# 		--proto_path=.


# generate3:
# 	protoc *.proto \
# 		--go_out=. \
# 		--go_opt=paths=source_relative \
# 		--go-grpc_out=. \
# 		--go-grpc_opt=paths=source_relative \
# 		--proto_path=.

helloworld:
	cd helloworld && \
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
    	helloworld.proto