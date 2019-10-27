FROM golang:1.13.0-buster

WORKDIR /go/src/app
COPY bin/main .
COPY certs/server.key .
COPY certs/server.crt .
ENV MICROSERVICE_KEY_FILE=/go/src/app/server.key 
ENV MICROSERVICE_CERT_FILE=/go/src/app/server.crt
ENV MICROSERVICE_SERVICE_ADDRESS="localhost:8080"

EXPOSE 8080

ENTRYPOINT /go/src/app/main
