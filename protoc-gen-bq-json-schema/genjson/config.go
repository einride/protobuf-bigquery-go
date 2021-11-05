package genjson

import "github.com/spf13/pflag"

type Config struct {
	Format bool
}

func (c *Config) AddToFlagSet(flags *pflag.FlagSet) {
	flags.BoolVar(
		&c.Format,
		"format",
		false,
		"Set to true to get a formatted json output.",
	)
}
