# -include .env

OUTPUT:=bin
EXECUTABLE:=proto2yaml
DOCKER_REGISTRY:=ghcr.io
DOCKER_USER:=krzko

VERSION:=v0.0.3
COMMIT=$(shell git describe --tags --always --long)
NOW=$(shell date +'%Y%m%d')

ARM5:=${EXECUTABLE}_${VERSION}_linux_arm_5
ARM6:=${EXECUTABLE}_${VERSION}_linux_arm_6
ARM7:=${EXECUTABLE}_${VERSION}_linux_arm_7
ARM8:=${EXECUTABLE}_${VERSION}_linux_arm64_8
DARWIN_AMD64:=${EXECUTABLE}_${VERSION}_darwin_amd64
DARWIN_ARM64:=${EXECUTABLE}_${VERSION}_darwin_arm64
FREEBSD:=${EXECUTABLE}_${VERSION}_freebsd_amd64
LINUX:=${EXECUTABLE}_${VERSION}_linux_amd64
WINDOWS:=${EXECUTABLE}_${VERSION}_windows_amd64

all: clean build

build: build-arm build-darwin-amd64 build-darwin-arm64 build-freebsd build-linux build-windows build-docker

build-no-docker: build-arm build-darwin-amd64 build-darwin-arm64 build-freebsd build-linux build-windows

build-arm:
	@echo "  🍒  Building binary ${ARM5}..."
	@[ -d "${OUTPUT}/${ARM5}" ] || mkdir -p "${OUTPUT}/${ARM5}"
	@env GOOS=linux GOARCH=arm GOARM=5 go build -o "${OUTPUT}/${ARM5}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${ARM5}.tar.gz" "${OUTPUT}/${ARM5}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${ARM5} complete"

	@echo "  🍒  Building binary ${ARM6}..."
	@[ -d "${OUTPUT}/${ARM6}" ] || mkdir -p "${OUTPUT}/${ARM6}"
	@env GOOS=linux GOARCH=arm GOARM=6 go build -o "${OUTPUT}/${ARM6}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${ARM6}.tar.gz" "${OUTPUT}/${ARM6}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${ARM6} complete"

	@echo "  🍒  Building binary ${ARM7}..."
	@[ -d "${OUTPUT}/${ARM7}" ] || mkdir -p "${OUTPUT}/${ARM7}"
	@env GOOS=linux GOARCH=arm GOARM=7 go build -o "${OUTPUT}/${ARM7}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${ARM7}.tar.gz" "${OUTPUT}/${ARM7}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${ARM7} complete"

	@echo "  🍒  Building binary${ARM8}..."
	@[ -d "${OUTPUT}/${ARM8}" ] || mkdir -p "${OUTPUT}/${ARM8}"
	@env GOOS=linux GOARCH=arm64 go build -o "${OUTPUT}/${ARM8}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${ARM8}.tar.gz" "${OUTPUT}/${ARM8}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${ARM8} complete"
	@echo

build-darwin-amd64:
	@echo "  🍎  Building binary ${DARWIN_AMD64}..."
	@[ -d "${OUTPUT}/${DARWIN_AMD64}" ] || mkdir -p "${OUTPUT}/${DARWIN_AMD64}"
	@env GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT}/${DARWIN_AMD64}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@zip -q "${OUTPUT}/${DARWIN_AMD64}".zip "${OUTPUT}/${DARWIN_AMD64}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${DARWIN_AMD64} complete"
	@echo

build-darwin-arm64:
	@echo "  🍏  Building binary ${DARWIN_ARM64}..."
	@[ -d "${OUTPUT}/${DARWIN_ARM64}" ] || mkdir -p "${OUTPUT}/${DARWIN_ARM64}"
	@env GOOS=darwin GOARCH=arm64 go build -o "${OUTPUT}/${DARWIN_ARM64}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@zip -q "${OUTPUT}/${DARWIN_ARM64}".zip "${OUTPUT}/${DARWIN_ARM64}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${DARWIN_ARM64} complete"
	@echo

build-freebsd:
	@echo "  👿  Building binary ${FREEBSD}..."
	@[ -d "${OUTPUT}/${FREEBSD}" ] || mkdir -p "${OUTPUT}/${FREEBSD}"
	@env GOOS=freebsd GOARCH=amd64 go build -o "${OUTPUT}/${FREEBSD}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${FREEBSD}.tar.gz" "${OUTPUT}/${FREEBSD}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${FREEBSD} complete"
	@echo

build-linux:
	@echo "  🐧  Building binary ${LINUX}..."
	@[ -d "${OUTPUT}/${LINUX}" ] || mkdir -p "${OUTPUT}/${LINUX}"
	@env GOOS=linux GOARCH=amd64 go build -o "${OUTPUT}/${LINUX}/${EXECUTABLE}" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@tar -czf "${OUTPUT}/${LINUX}.tar.gz" "${OUTPUT}/${LINUX}/${EXECUTABLE}"
	@echo "  ✨  Build binary ${LINUX} complete"
	@echo

build-windows:
	@echo "  💾  Building binary ${WINDOWS}..."
	@[ -d "${OUTPUT}/${WINDOWS}" ] || mkdir -p "${OUTPUT}/${WINDOWS}"
	@env GOOS=windows GOARCH=amd64 go build -o "${OUTPUT}/${WINDOWS}/${EXECUTABLE}.exe" -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.commit=${COMMIT} -X main.date=${NOW}" cmd/${EXECUTABLE}/main.go
	@zip -q "${OUTPUT}/${WINDOWS}.zip" "${OUTPUT}/${WINDOWS}/${EXECUTABLE}.exe"
	@echo "  ✨  Build binary ${WINDOWS} complete"
	@echo

build-docker:
	@echo "  🐳  Building docker image ${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION}..."
	@docker image build --build-arg "COMMIT=${COMMIT}" --build-arg "NOW=${NOW}" --build-arg "VERSION=${VERSION}" -t "${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION}" .
	@echo "  ✨  Build docker image ${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION} complete"
	@echo

clean:
	@echo "  🧹  Cleaning build artifacts..."
	@rm -rf $(OUTPUT) 2> /dev/null
	@rm -r *.txt 2> /dev/null || true
	@docker image prune -f
	@echo "  ✨  Cleaning build artifacts complete"

fmt:
	go fmt ./...

push-docker:
	@echo "  🐳  Pushing docker image ${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION}..."
	@docker push "${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION}"
	@echo "  ✨  Docker image ${DOCKER_REGISTRY}/${DOCKER_USER}/${EXECUTABLE}:${VERSION} pushed"
	@echo

run:
	go run cmd/${EXECUTABLE}/main.go

test:
	go test -v ./...

default: all
