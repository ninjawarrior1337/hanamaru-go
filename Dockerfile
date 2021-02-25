FROM golang as builder
WORKDIR /hanamaru
COPY . .
RUN go generate
ENV CGO_ENABLED=0
RUN go build -ldflags='-s -w' -tags="ij,jp"

FROM alpine
ENV IN_DOCKER=true
RUN apk add --no-cache youtube-dl ffmpeg
WORKDIR /app
VOLUME [ "/data" ]
COPY --from=builder /hanamaru/hanamaru-go .
CMD [ "/app/hanamaru-go" ]