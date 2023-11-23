PROTO_FILE := grpc/server.proto
PROTO_TIME_FILE := time/time.proto
OUT_DIR := .

# To run `make generate`
.PHONY: generate
.PHONY: generate-time
generate:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_FILE)
generate-time:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_TIME_FILE)
