package config

type EVMAPIConfig struct {
	// Enabled is a flag that indicates whether the provider is API based.
	Enabled bool `mapstructure:"enabled" toml:"enabled"`
}

// ValidateBasic performs basic validation of the API config.
func (c *EVMAPIConfig) ValidateBasic() error {
	return nil
}
