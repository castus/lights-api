FROM --platform=$BUILDPLATFORM golang:1.23 AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETARCH
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

WORKDIR /data
COPY . /data
RUN GOOS=linux GOARCH=$TARGETARCH go build -o lights-api

FROM scratch
COPY --from=builder /data/lights-api ./
CMD ["./lights-api"]
