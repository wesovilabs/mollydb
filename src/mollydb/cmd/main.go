package main

import (
	"log"
	"mollydb/pkg/api"
	"mollydb/pkg/config"
	"net/http"
)

func main() {
	log.Printf("Launching mollydb on %s", config.MollyConfig.Server.Address)
	http.Handle(api.Handler())
	fs := http.FileServer(http.Dir(config.MollyConfig.Server.Graphiql))
	http.Handle("/", fs)
	http.ListenAndServe(config.MollyConfig.Server.Address, nil)
}
