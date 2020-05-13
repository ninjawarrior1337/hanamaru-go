FROM golang AS builder
WORKDIR /mage
ENV CI true
RUN git clone https://github.com/magefile/mage && cd mage && go run bootstrap.go
WORKDIR /app
COPY ./ /app/
RUN mage buildDocker

FROM alpine
WORKDIR /app
COPY --from=builder /app/hanamaru .
CMD ["./hanamaru"]