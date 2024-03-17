FROM golang as builder
WORKDIR /hanamaru

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go generate
RUN go build -ldflags='-s -w' -tags="ij,jp"

FROM registry.access.redhat.com/ubi9/ubi-minimal
ENV IN_DOCKER=true
WORKDIR /app
VOLUME [ "/data" ]
COPY --from=builder /hanamaru/hanamaru-go .
CMD [ "/app/hanamaru-go" ]