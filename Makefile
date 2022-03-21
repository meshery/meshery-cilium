GOPATH = $(shell go env GOPATH)

check: error
	golangci-lint run

check-clean-cache:
	golangci-lint cache clean

protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/meshery/meshery/master/meshes/meshops.proto

proto:
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

docker:
	docker build -t layer5/meshery-cilium .

docker-run:
	(docker rm -f meshery-cilium) || true
	docker run --name meshery-cilium -d \
	-p 10012:10012 \
	-e DEBUG=true \
	layer5/meshery-cilium

run:
	go mod tidy; \
	DEBUG=true go run main.go

.PHONY: error
error:
	go run github.com/meshery/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers

.PHONY: local-check
local-check: tidy
local-check: golangci-lint

.PHONY: tidy
tidy:
	@echo "Executing go mod tidy"
	go mod tidy

.PHONY: golangci-lint
golangci-lint: $(GOLANGCILINT)
	@echo
	$(GOPATH)/bin/golangci-lint run

$(GOLANGCILINT):
	(cd /; GO111MODULE=on GOPROXY="direct" GOSUMDB=off go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.30.0)
