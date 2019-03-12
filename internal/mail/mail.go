package mail

import (
	"log"

	"github.com/spf13/viper"
)

var (
	factories map[string]Factory
)

func init() {
	factories = make(map[string]Factory)
	viper.SetDefault("mailer", "null")
}

// Register is called by external implementations of the mailer
// interface to be made available in the registry.
func Register(name string, f Factory) {
	if _, ok := factories[name]; ok {
		// Already registered...
		return
	}
	log.Printf("Registering mail implementation %s", name)
	factories[name] = f
}

// Initialize invokes the factory specified by the viper token
// "mailer" and returns the result of this call.
func Initialize() (Mailer, error) {
	if f, ok := factories[viper.GetString("mailer")]; ok {
		return f()
	}
	return nil, ErrUnknownMailer
}
