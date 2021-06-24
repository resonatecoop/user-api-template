# Build stage
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/resonatecoop/user-api
RUN cd /build && git clone https://github.com/resonatecoop/user-api

RUN cd user-api && go build

EXPOSE 11000

ENTRYPOINT ["./docker-entrypoint.sh"]
