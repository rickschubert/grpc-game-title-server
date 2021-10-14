compile-proto:
	mkdir -p grpc_interfaces && \
    protoc --go_out=./grpc_interfaces --go_opt=paths=source_relative \
    --go-grpc_out=./grpc_interfaces --go-grpc_opt=paths=source_relative \
    gaming_stats.proto
