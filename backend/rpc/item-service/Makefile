.SILENT:

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lucamatsumoto/go-sales-manager/backend/rpc/item-service \
	proto/item/item.proto 

	docker build -t item-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDR=:50051 -e MICRO_REGISTRY=mdns item-service

