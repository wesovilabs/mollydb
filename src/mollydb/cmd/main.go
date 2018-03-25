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
		database.Molly = &database.Storage{
			Root: &model.Folder{
				Name:      "root",
				Documents: make(map[string]*model.Document),
				Folders:   make(map[string]*model.Folder),
			},
			Path: config.MollyConfig.Storage.LocalPath,
		}
		database.Molly.SetUp()
		go database.Molly.Watch()
	}()
	http.Handle(api.Handler())
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(config.MollyConfig.Server.Address, nil)
}
