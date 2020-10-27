FROM golang:1.15.3-buster as build
ENV BACKEND_HOST=localhost
ENV BACKEND_PORT=56789
ENV FRONTEND_HOST=localhost
ENV FRONTEND_PORT=12345

EXPOSE 12345
EXPOSE 56789
RUN apt update && apt upgrade -y
RUN apt install libzmq3-dev pkg-config wget tar -y
WORKDIR /go/src/certificates
COPY ./certs/certs-dec-2019.tar .
RUN wget https://raw.githubifc.iad.ca.inet/rplhy997/AIengineer/master/certificates/install-certificates-linux.sh?token=AAAAP467Y4IG6LJ5F7Z3CRK7UFP5Q -O install-certificates-linux.sh --no-check-certificate
RUN sed -i 's|/usr/share/ca-certificates|/etc/ssl/certs|g' install-certificates-linux.sh
RUN ls
RUN chmod +x install-certificates-linux.sh
RUN ./install-certificates-linux.sh
WORKDIR /go/src/client
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["efog-iotedge-hub-clientsimulator"]