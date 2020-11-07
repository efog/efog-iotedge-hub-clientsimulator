FROM golang:1.15.3-buster as build
ENV PUBLISHER_ENDPOINT=tcp://localhost:7000
ENV SUBSCRIBER_ENDPOINT=tcp://localhost:7001

RUN apt update && apt upgrade -y
RUN apt install libzmq3-dev pkg-config wget tar -ypa
WORKDIR /go/src/client
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["efog-iotedge-hub-clientsimulator"]