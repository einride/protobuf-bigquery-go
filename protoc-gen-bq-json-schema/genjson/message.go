package genjson

import (
	"fmt"

	"go.einride.tech/protobuf-bigquery/encoding/jsonbq"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateSchemaFile(gen *protogen.Plugin, msg *protogen.Message, config Config) error {
	g := gen.NewGeneratedFile(fmt.Sprintf("%s.json", msg.Desc.Name()), "")
	schema := protobq.SchemaOptions{
		UseModeFromFieldBehavior: config.UseModeFromFieldBehavior,
	}.InferMessageSchema(msg.Desc)
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
