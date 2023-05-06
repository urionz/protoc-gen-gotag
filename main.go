package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/urionz/protoc-gen-gotag/module"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(pgs.DebugEnv("GOTAG_DEBUG"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
