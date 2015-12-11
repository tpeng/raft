COVERPROFILE=cover.out

default: test

cover:
	go test -coverprofile=$(COVERPROFILE) .
	go tool cover -html=$(COVERPROFILE)
	rm $(COVERPROFILE)

dependencies:
	go get -d .

test:
	go test -i ./...
	go test -v ./...

proto:
	protoc ./proto/raft.proto --go_out=plugins=grpc:.

.PHONY: coverage dependencies test
