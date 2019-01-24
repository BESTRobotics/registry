package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/http"
)

func main() {
	pflag.String("core.bind", "", "Address to bind to")
	pflag.Int("core.port", 8080, "Port to bind to")
	pflag.String("dev.webroot", "web", "Webroot during development")
	pflag.Bool("dev.extweb", true, "Use local webroot")

	viper.BindPFlags(pflag.CommandLine)

	http.Serve()
}
