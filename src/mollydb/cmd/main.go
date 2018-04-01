package main

import (
	"log"
	"mollydb/pkg/api"
	"mollydb/pkg/config"
	"net/http"
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
	for _, s := range config.MollyConfig.StorageList {
		storage:=&model.StorageList{
			Name:      s.Name,
			Path:      s.Path,
			Documents: make(map[string]*model.Document),
		}
		database.Molly.SetUpStorage(storage)
	}
	**/
}
