package main

import (
	"net/http"
	"mollydb/pkg/api"
	"mollydb/pkg/config"
	"mollydb/pkg/model"
	"mollydb/pkg/database"
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
	database.Molly = &database.Database{
		StorageDict:   make(map[string]*model.Storage),
		PropertyCache: make(map[string]*model.Property),
		PropertyHooks: make(map[string][]database.Hook),
	}
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
