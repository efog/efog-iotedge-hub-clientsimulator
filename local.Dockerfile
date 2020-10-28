FROM golang:1.15.3-buster as build
ENV BACKEND_HOST=localhost
ENV BACKEND_PORT=56789
ENV FRONTEND_HOST=localhost
ENV FRONTEND_PORT=12345

EXPOSE 12345
EXPOSE 56789
RUN apt update && apt upgrade -y
RUN apt install libzmq3-dev pkg-config wget tar -ypa
WORKDIR /go/src/client
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["efog-iotedge-hub-clientsimulator"]