
.PHONY: gensc protoc-gen-go-gensc-server protoc-gen-go-gensc-client

all: gensc protoc-gen-go-gensc-server protoc-gen-go-gensc-client

gensc:
	go build -o gensc cmd/gensc/main.go


protoc-gen-go-gensc-server:
	go build -o protoc-gen-go-gensc-server cmd/protoc-gen-go-gensc-server/main.go

protoc-gen-go-gensc-client:
	go build -o protoc-gen-go-gensc-client cmd/protoc-gen-go-gensc-client/main.go

