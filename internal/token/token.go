package token

import (
	"time"

	"github.com/spf13/viper"
)

// The Config struct contains information that should be used when
// generating a token.
type Config struct {
	Lifetime  time.Duration
	Issuer    string
	IssuedAt  time.Time
	NotBefore time.Time
}

// GetConfig returns a struct containing the configuration for the
// token service to use while issuing tokens.
func GetConfig() Config {
	return Config{
		Lifetime:  viper.GetDuration("token.lifetime"),
		IssuedAt:  time.Now(),
		NotBefore: time.Now(),
	}
}
