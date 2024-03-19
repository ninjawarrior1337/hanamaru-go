FROM rust as builder_rs
WORKDIR /hanamaru
COPY ./lib ./
RUN cargo build --release

FROM golang as builder_go
WORKDIR /hanamaru

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN mkdir -p ./lib/target/release/
COPY --from=builder_rs /hanamaru/target/release/libhanamaru_lib.a  ./lib/target/release/
RUN go generate
RUN go build -ldflags='-s -w' -tags="ij,jp"

FROM debian:12-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENV IN_DOCKER=true
WORKDIR /app
VOLUME [ "/data" ]
COPY --from=builder_go /hanamaru/hanamaru-go .
CMD [ "/app/hanamaru-go" ]