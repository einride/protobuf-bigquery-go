package main

import (
	"github.com/spf13/pflag"
	"go.einride.tech/protobuf-bigquery/protoc-gen-bq-json-schema/genjson"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flagSet := pflag.NewFlagSet("protoc-gen-bq-json-schema", pflag.ContinueOnError)
	var config genjson.Config
	config.AddToFlagSet(flagSet)
	protogen.Options{ParamFunc: flagSet.Set}.Run(func(gen *protogen.Plugin) error {
		return genjson.Run(gen, config)
	})
}
