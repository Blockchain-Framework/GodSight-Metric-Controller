FROM golang:1.18-alpine AS builder

WORKDIR /src/
COPY . .
ARG RELEASE
ARG COMMIT
RUN VERSION_PKG=$(go list -m)/pkg && \
    BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S') && \
    CGO_ENABLED=0 GO111MODULE=on \
	go build \
		-ldflags "-w -s \
		-X $VERSION_PKG/version.Release=$RELEASE -X ${VERSION_PKG}/version.Commit=${COMMIT} -X ${VERSION_PKG}/version.BuildTime=${BUILD_TIME}" \
		-a -o /bin/serviceaggregator \
		./cmd/main.go
RUN apk --no-cache add ca-certificates && update-ca-certificates
FROM scratch
COPY --from=builder /bin/serviceaggregator /bin/serviceaggregator
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./bin/serviceaggregator"]
