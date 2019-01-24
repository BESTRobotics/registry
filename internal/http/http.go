package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	http.HandleFunc("/status", okHandler)
}

// Serve is the entrypoint to the http package that serves the API.
func Serve() {
	if viper.GetBool("dev.extweb") {
		log.Println("Using external webroot", viper.GetString("dev.webroot"))
		http.Handle("/", http.FileServer(http.Dir(viper.GetString("dev.webroot"))))
	}

	bindstr := fmt.Sprintf("%s:%d",
		viper.GetString("core.bind"),
		viper.GetInt("core.port"))

	log.Fatal(http.ListenAndServe(bindstr, nil))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("System OK"))
}
