

all: gen-server protoc-gen-go-gensc-server protoc-gen-go-gensc-client

gen-server:
	go build -o gen-server cmd/gen-server/main.go


protoc-gen-go-gensc-server:
	go build -o protoc-gen-go-gensc-server cmd/protoc-gen-go-gensc-server/main.go

protoc-gen-go-gensc-client:
	go build -o protoc-gen-go-gensc-client cmd/protoc-gen-go-gensc-client/main.go

