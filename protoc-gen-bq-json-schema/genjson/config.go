package genjson

import "github.com/spf13/pflag"

type Config struct {
	Format                   bool
	UseModeFromFieldBehavior bool
}

func (c *Config) AddToFlagSet(flags *pflag.FlagSet) {
	flags.BoolVar(
		&c.Format,
		"format",
		false,
		"Set to true to get a formatted json output.",
	)
	flags.BoolVar(
		&c.UseModeFromFieldBehavior,
		"field_behavior_mode",
		false,
		"Set to true to use google api field behavior for setting field mode.",
	)
}
