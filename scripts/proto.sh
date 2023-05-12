protoc --go_out=grpc/gen --go_opt=paths=source_relative \
  --go-grpc_out=grpc/gen --go-grpc_opt=paths=source_relative \
  grpc/proto/v1/bus/bus.proto
