package main

import (
	"flag"
	"fmt"

	"github.com/lujingwei002/gensc/gen/gen_grpc_client"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

// from https://github.com/grpc/grpc-go/blob/cmd/protoc-gen-go-grpc/v1.3.0/cmd/protoc-gen-go-grpc/main.go

const version = "1.3.0"

var opts gen_grpc_client.Options

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-gclient %v\n", version)
		return
	}

	var flags flag.FlagSet
	opts.RequireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")
	opts.ProtoAndGrpcPackage = flags.String("proto_and_grpc_pacakge", "", "proto and grpc package")
	opts.Version = version
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			// log.Println(f)
			gen_grpc_client.GenerateFile(gen, f, opts)
		}
		return nil
	})
}
