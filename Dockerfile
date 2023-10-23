FROM golang:1.21.3 AS builder

WORKDIR /data
COPY . /data
RUN GOOS=linux GOARCH=amd64 go build -o lights-api

FROM ubuntu:mantic
COPY --from=builder /data/lights-api ./
CMD ["./lights-api"]
