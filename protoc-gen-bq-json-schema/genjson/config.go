package genjson

import (
	"flag"

	"go.einride.tech/protobuf-bigquery/encoding/protobq"
)

type Config struct {
	Format        bool
	SchemaOptions protobq.SchemaOptions
}

func (c *Config) AddToFlagSet(flags *flag.FlagSet) {
	flags.BoolVar(
		&c.Format,
		"format",
		false,
		"Set to true to get a formatted json output.",
	)
	flags.BoolVar(
		&c.SchemaOptions.UseEnumNumbers,
		"enum_numbers",
		false,
		"Map proto enum fields to BQ INTEGER fields.",
	)
	flags.BoolVar(
		&c.SchemaOptions.UseDateTimeWithoutOffset,
		"datetime_without_offset",
		false,
		"Convert google.type.DateTime fields to DATETIME and discard the optional time offset.",
	)
	flags.BoolVar(
		&c.SchemaOptions.UseOneofFields,
		"oneof_fields",
		false,
		"Add an extra STRING field for oneof fields with the name of the oneof, "+
			"to contain the name of the field that is set.",
	)
	flags.BoolVar(
		&c.SchemaOptions.UseModeFromFieldBehavior,
		"mode_from_field_behavior",
		false,
		"Set the mode of a field to REQUIRED if the field is defined with REQUIRED behavior in proto.",
	)
}
