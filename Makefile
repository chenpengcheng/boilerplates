VERSION := $(shell git describe --tags --always --dirty="-dev")
LDFLAGS := -ldflags='-X "main.version=$(VERSION)"'

Q=@

dep:
	$Qdep ensure -v

clean:
	$Qrm -rf vendor/ && dep ensure -v

run: dep
	$Qgo run cmd/service/main.go

vet: dep
	$Qgo vet ./...

test: dep
	$Qgo test -v -count=1 -race ./...

build:
	$Qdocker build --build-arg VERSION=$(VERSION) \
		-t pengchengchen/rpc-server:$(VERSION) \
		-t pengchengchen/rpc-server:latest .

docker:
	$Qdocker-compose up -d

teardown:
	$Qdocker-compose down -v
	$Q rm -fr data/mysql data/mongodb

release:
	$Qdocker push pengchengchen/rpc-server:$(VERSION)
	$Qdocker push pengchengchen/rpc-server:latest

.PHONY: dep clean run vet test build docker teardown release
