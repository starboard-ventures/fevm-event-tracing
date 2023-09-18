COMMIT_ID = $(shell git rev-parse --short HEAD)
ifeq ($(COMMIT_ID),)
COMMIT_ID = 'latest'
endif

.PHONY: test
IMAGE_PREFIX ?= 129862287110.dkr.ecr.us-east-2.amazonaws.com/infra
REGISTRY_SERVER ?= 129862287110.dkr.ecr.us-east-2.amazonaws.com/

help:
	@echo
	@echo "  binary - build binary"
	@echo "  build-fevm-event - build docker images for apline"
	@echo "  swag - regenerate swag"
	@echo "  build-all - build docker images for apline"
	@echo "  push images to docker hub"

swag:
	swag init -g cmd/busi/main.go

binary:
	go build -o bin/fevm-event cmd/busi/main.go

test:
	go clean -testcache
	gotestsum --format pkgname

build-fevm-event:
	docker build -t $(IMAGE_PREFIX)/fevm-event:$(COMMIT_ID) -f build/Dockerfile .

build-all: build-fevm-event

push:
	docker push $(IMAGE_PREFIX)/fevm-event:$(COMMIT_ID)
