# Build stage
ARG RELEASE_TAG=master
FROM golang:latest

ARG RELEASE_TAG

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN apt-get -y update
RUN apt-get install -y libpq-dev postgresql-client
RUN git clone --branch ${RELEASE_TAG} --single-branch --depth 1 https://github.com/resonatecoop/user-api

WORKDIR /build/user-api

RUN make install
RUN git submodule update --init
RUN make generate
RUN go build

EXPOSE 11000

ENTRYPOINT ["sh", "docker-entrypoint.sh"]
