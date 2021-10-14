.PHONY: compile-proto

compile-proto:
    cd gamingstats && \
    protoc --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    gaming_stats.proto && \
    cd ../
