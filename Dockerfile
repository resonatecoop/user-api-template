# Build stage
FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN apt-get -y update
RUN apt-get install -y libpq-dev postgresql-client
RUN go get github.com/resonatecoop/user-api
RUN cd /build && git clone https://github.com/resonatecoop/user-api

RUN cd user-api && go build

EXPOSE 11000

WORKDIR /build/user-api

ENTRYPOINT ["sh", "docker-entrypoint.sh"]
