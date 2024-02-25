BIN := serviceaggregator
RELEASE := 0.0.1
IMAGE := app-aggregator

COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

APPLICATION_MODULE := $(shell go list -m)
VERSION_PKG := ${APPLICATION_MODULE}"/pkg"

AIR_CMD = ${HOME}/go/bin/air
BIN_DIR = bin
OUT_BIN := ${BIN_DIR}/${OS}_${ARCH}/${BIN}

# default target
all: build

# clean generated artifacts
clean:
	@echo "---- Cleaning Service Aggregator Build ----"
	rm -rf ${BIN_DIR}

# generate binary locally
build: clean
	@echo "---- Building Controller (Local) ----"
	CGO_ENABLED=0 GO111MODULE=on \
	go build \
		-ldflags "-w -s \
		-X ${VERSION_PKG}/version.Release=${RELEASE} -X ${VERSION_PKG}/version.Commit=${COMMIT} -X ${VERSION_PKG}/version.BuildTime=${BUILD_TIME}" \
		-a -o ${OUT_BIN} \
		./cmd/main.go

# run application using generated binary
run: build
	@echo "---- Running Controller using Binary ----"
	./${OUT_BIN}

# run application using sources
run-dev:
	@echo "---- Running Controller using Source ----"
	@go run ./cmd/main.go

#run swagger doc generation
swagger:
	@echo "---- Running SWAG, Generating the swagger documentation ----"
	@swag init -g cmd/app_deployer/main.go

vet:
	@go vet ./...

# run application with live reload
dev:
	@echo "---- Running Controller With Live Reload ----"
	@$(AIR_CMD) -c air.toml

container:
	@echo "---- Creating ${IMAGE}:latest ----"
	@docker build \
		--no-cache \
		-t $(IMAGE):latest \
		--build-arg RELEASE=${RELEASE} --build-arg COMMIT=${COMMIT} \
		.

run-image: container
	@echo "---- Running ${IMAGE}:$(shell docker images -q $(IMAGE)) ----"
	@docker-compose up