package main

import (
	"net/http"
	"mollydb/pkg/api"
	"mollydb/pkg/config"
	"mollydb/pkg/model"
	"mollydb/pkg/database"
)

func main() {
	println("Launching mollydb")
	go func() {
		storage := &model.Storage{
			Name:      "root",
			Path:      config.MollyConfig.Storage.LocalPath,
			Documents: make(map[string]*model.Document),
		}
		database.Molly = &database.Database{
			StorageDict: map[string]*model.Storage{
				"root": storage,
			},
		}
		database.Molly.SetUp()
		go database.Molly.Watch()
	}()
	http.Handle(api.Handler())
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(config.MollyConfig.Server.Address, nil)
}
