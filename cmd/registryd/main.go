package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/db"
	"github.com/BESTRobotics/registry/internal/http"
	"github.com/BESTRobotics/registry/internal/mechgreg"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	pflag.String("http.bind", "", "Address to bind to")
	pflag.Int("http.port", 8080, "Port to bind to")
	pflag.String("dev.webroot", "web", "Webroot during development")
	pflag.Bool("dev.extweb", true, "Use local webroot")
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	log.Println("regitryd initializing")
	log.Println("Preparing to serve")
	db, err := db.Open()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	rb := mechgreg.ResourceBundle{
		DB: db,
	}

	mg, err := mechgreg.New(&rb)
	if err != nil {
		log.Panic(err)
	}

	s, err := http.New(mg)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Launching network services")
	s.Serve()
}
