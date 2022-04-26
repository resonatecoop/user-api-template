# Build stage
ARG RELEASE_TAG=master
FROM golang:latest

ARG RELEASE_TAG

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN apt-get -y update
RUN apt-get install -y libpq-dev postgresql-client
RUN go install github.com/resonatecoop/user-api@${RELEASE_TAG}
RUN cd /build && git clone --branch ${RELEASE_TAG} --single-branch --depth 1 https://github.com/resonatecoop/user-api

RUN cd user-api && go build

EXPOSE 11000

WORKDIR /build/user-api

ENTRYPOINT ["sh", "docker-entrypoint.sh"]
