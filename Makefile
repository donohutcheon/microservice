build:
	go build -o bin/main main.go

run:
	go run main.go

docker: build
	docker build -t donohutcheon/microservice:1.0 .
	
