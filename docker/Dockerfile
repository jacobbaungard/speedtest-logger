ARG GOVER=1.21.3

FROM golang:$GOVER-bullseye AS builder

ARG LIBRESPEED_VER=1.0.10

WORKDIR /tmp/
RUN curl -L https://github.com/librespeed/speedtest-cli/releases/download/v"$LIBRESPEED_VER"/librespeed-cli_"$LIBRESPEED_VER"_linux_amd64.tar.gz -o speedtest-cli.tar.gz && \
  tar -xf /tmp/speedtest-cli.tar.gz

WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./*.go ./
COPY ./internal/ ./internal/

RUN go build

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/speedtest-logger ./speedtest-logger
COPY --from=builder /tmp/librespeed-cli /usr/bin/librespeed-cli
COPY ./docker/init.sh /app/

# ensure we have trusted certs installed
RUN apt-get update && apt-get --no-install-recommends -y install ca-certificates && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/app/init.sh"]
CMD ["/app/speedtest-logger"]
