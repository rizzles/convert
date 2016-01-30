package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mandibleConf "./config"
	processors "./imageprocessor"
	mandible "./server"
)

func main() {
	configFile := "config/conf.json" //os.Getenv("MANDIBLE_CONF")

	config := mandibleConf.NewConfiguration(configFile)

	var server *mandible.Server

	if os.Getenv("AUTHENTICATION_HMAC_KEY") != "" {
		key := []byte(os.Getenv("AUTHENTICATION_HMAC_KEY"))
		auth := mandible.NewHMACAuthenticatorSHA256(key)
		server = mandible.NewAuthenticatedServer(config, processors.EverythingStrategy, auth)
	} else {
		server = mandible.NewServer(config, processors.EverythingStrategy)
	}
	muxer := http.NewServeMux()
	server.Configure(muxer)

	port := fmt.Sprintf(":%d", server.Config.Port)

	log.Printf("Listening on Port: %s", port)

	http.ListenAndServe(port, muxer)
}
