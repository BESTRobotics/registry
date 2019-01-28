package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/db"
	"github.com/BESTRobotics/registry/internal/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	log.Println("regitryd initializing")
	pflag.String("core.bind", "", "Address to bind to")
	pflag.Int("core.port", 8080, "Port to bind to")
	pflag.String("dev.webroot", "web", "Webroot during development")
	pflag.Bool("dev.extweb", true, "Use local webroot")

	viper.BindPFlags(pflag.CommandLine)

	log.Println("Preparing to serve")
	db, err := db.Open()
	if err != nil {
		return
	}
	db.Close()

	log.Println("Launching network services")
	http.Serve()
}
