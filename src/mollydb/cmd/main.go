package main

import (
	"net/http"
	"mollydb/pkg/api"
	"mollydb/pkg/config"
	"log"
)

func main() {
	log.Printf("Launching mollydb on %s", config.MollyConfig.Server.Address)
	go loadMollyDB()
	http.Handle(api.Handler())
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(config.MollyConfig.Server.Address, nil)
}

func loadMollyDB() {

	/**
	for _, s := range config.MollyConfig.Storages {
		storage:=&model.Storage{
			Name:      s.Name,
			Path:      s.Path,
			Documents: make(map[string]*model.Document),
		}
		database.Molly.SetUpStorage(storage)
	}
	**/
}
