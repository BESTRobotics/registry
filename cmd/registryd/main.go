package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/http"
	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/mechgreg"
	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"

	_ "github.com/BESTRobotics/registry/internal/mail/null"
)

func init() {
	pflag.String("core.home", ".", "Home directory")
	pflag.Int("core.bootstrap", 0, "Grant superadmin to the named userID")
	pflag.String("http.bind", "", "Address to bind to")
	pflag.Int("http.port", 8080, "Port to bind to")
	pflag.String("dev.webroot", "web/build", "Webroot during development")
	pflag.Bool("dev.extweb", false, "Use local webroot")
	pflag.Bool("dev.cors", false, "Add CORS header for *")
	pflag.String("storage.path", ".", "Path to the data area")
	pflag.Bool("token.generate", false, "Enable generation of token keys")
	pflag.Int("token.bits", 2048, "How many bits to include in token keys")
	pflag.Bool("internal.reindex", false, "Reindex internal data structures")
}

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	log.Println(viper.GetBool("dev.cors"))
	log.Println("regitryd initializing")
	log.Println("Preparing to serve")

	dbPath := filepath.Join(viper.GetString("storage.root"), "registry.db")

	st, err := storm.Open(dbPath)
	if err != nil {
		log.Panic(err)
	}

	mailer, err := mail.Initialize()
	if err != nil {
		log.Panic(err)
	}

	tkn, err := token.NewRSA()
	if err != nil {
		log.Panic(err)
	}

	rb := mechgreg.ResourceBundle{
		StormDB: st,
		Mailer:  mailer,
	}

	mg, err := mechgreg.New(rb)
	if err != nil {
		log.Panic(err)
	}

	if viper.GetInt("core.bootstrap") != 0 {
		user, err := mg.GetUser(viper.GetInt("core.bootstrap"))
		if err != nil {
			log.Panic("Bootstrap load error:", err)
		}
		user.GrantCapability(models.CapSuperAdmin)
		if err := mg.ModUser(user); err != nil {
			log.Panic("Bootstrap save error:", err)
		}
		log.Println("Bootstrap Successful")
		os.Exit(0)
	}

	s, err := http.New(mg, tkn, mailer)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Launching network services")
	s.Serve()
}
