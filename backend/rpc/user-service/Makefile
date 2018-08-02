.SILENT:

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service \
	proto/user/user.proto 
	
	docker build -t user-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDR=:50051 -e MICRO_REGISTRY=mdns user-service

