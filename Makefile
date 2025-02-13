# Определяем переменные
PROTO_DIR = proto
OUTPUT_DIR = ./gen/go
PROTOC = protoc
PROTOC_GEN_GO = --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative
PROTOC_GEN_GRPC = --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

all: generate

generate:
	$(PROTOC) -I $(PROTO_DIR) $(PROTO_DIR)/sso/*.proto $(PROTOC_GEN_GO) $(PROTOC_GEN_GRPC)

gen: generate
