up:
	docker-compose up

down:
	docker-compose down --remove-orphans

.PHONY: consumer
consumer:
	go run consumer/main.go

message:
	go run producer/main.go
