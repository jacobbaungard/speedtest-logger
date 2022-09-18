FROM golang:1.19.1-bullseye AS builder

WORKDIR /tmp/
# TODO: Set version as a variable instead...
RUN curl -L https://github.com/librespeed/speedtest-cli/releases/download/v1.0.10/librespeed-cli_1.0.10_linux_amd64.tar.gz -o speedtest-cli.tar.gz
RUN ls /tmp/
RUN tar -xf /tmp/speedtest-cli.tar.gz
RUN ls /tmp/

WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./*.go ./
COPY ./internal/ ./internal/

RUN go build

FROM debian:bullseye
WORKDIR /app
COPY --from=0 /app/speedtest-logger ./speedtest-logger
COPY --from=0 /tmp/librespeed-cli /usr/bin/librespeed-cli
COPY ./docker/init.sh /app/

# ensure we have trusted certs installed
RUN apt-get update && apt-get -y install ca-certificates

ENTRYPOINT ["/app/init.sh"]
CMD ["/app/speedtest-logger"]
