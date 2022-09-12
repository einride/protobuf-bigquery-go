package genjson

import (
	"fmt"
	"strings"

	"go.einride.tech/protobuf-bigquery/encoding/jsonbq"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateSchemaFile(gen *protogen.Plugin, msg *protogen.Message, config Config) error {
	filename := strings.ReplaceAll(strings.ToLower(string(msg.Desc.FullName())), ".", "_")
	g := gen.NewGeneratedFile(fmt.Sprintf("%s.json", filename), "")
	schema := config.SchemaOptions.InferMessageSchema(msg.Desc)
	jsonOpt := jsonbq.MarshalOptions{}
	if config.Format {
		jsonOpt.Indent = "  "
	}
	out, err := jsonOpt.MarshalSchema(schema)
	if err != nil {
		return err
	}
	g.P(string(out))
	return nil
}
