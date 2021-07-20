FROM golang:1.16.6-alpine3.13 AS builder
WORKDIR /src/
COPY . /src/
ARG COMMIT
ARG NOW
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/proto2yaml -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/proto2yaml/main.go

FROM scratch
ARG COMMIT
ARG NOW
ARG VERSION
LABEL maintainer="Kristof Kowalski <k@ko.wal.ski>" \
    org.opencontainers.image.title="proto2yaml" \
    org.opencontainers.image.description="A command-line utility to export Protocol Buffers (proto) files to YAML, and JSON" \
    org.opencontainers.image.authors="Kristof Kowalski <k@ko.wal.ski>" \
    org.opencontainers.image.vendor="Kristof Kowalski" \
    org.opencontainers.image.documentation="https://github.com/krzko/proto2yaml/docs" \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.version=$VERSION \
    org.opencontainers.image.url="https://ko.wal.ski" \
    org.opencontainers.image.source="https://github.com/krzko/proto2yaml.git" \
    org.opencontainers.image.revision=$COMMIT \
    org.opencontainers.image.created=$NOW
COPY --from=builder /bin/proto2yaml /bin/proto2yaml
ENTRYPOINT ["/bin/proto2yaml"]