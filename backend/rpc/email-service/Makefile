build:
	docker build -t email-service .

run:
	docker run --net="host" -p 50054 -e MICRO_SERVER_ADDR=:50054 -e MICRO_REGISTRY=mdns email-service
		