generate_proto:
	protoc \
	--go-grpc_out=proto \
	--go-grpc_opt=paths=source_relative \
	proto/pong.proto
