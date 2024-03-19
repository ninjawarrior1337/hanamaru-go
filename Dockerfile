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

FROM ubuntu
RUN apt update && apt install -y ca-certificates ffmpeg
ENV IN_DOCKER=true
WORKDIR /app
VOLUME [ "/data" ]
COPY --from=builder_go /hanamaru/hanamaru-go .
CMD [ "/app/hanamaru-go" ]