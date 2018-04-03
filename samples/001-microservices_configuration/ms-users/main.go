package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"strings"
)

const (
	mollydb = "http://mollydb:9090/graphql"
	query   = `query properties {
        properties(storage: "ms", document:"ms-users") {
            path
            key
            value
          }
    }`
)

func init() {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	body := strings.NewReader(query)
	req, err := http.NewRequest(http.MethodPost, mollydb, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/graphql")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	response := &MollyDBResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		// handle err
	}
	cfg = &configuration{}
	cfg.load(response.Data.Properties)
}

type MollyDBResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Path  string `json:"path"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func setUpConfiguration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating configuration")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var property Property
	err = json.Unmarshal(b, &property)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Property %v", property)
	cfg.load([]Property{property})
	fmt.Println(property)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/configuration", setUpConfiguration)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
	})
	http.ListenAndServe(":3333", r)
}

var cfg *configuration

type configuration struct {
	logLevel string
	dbURI    string
}

func (cfg *configuration) load(properties []Property) {
	for _, p := range properties {
		switch p.Key {
		case "logLevel":
			cfg.logLevel = p.Value
			lvl, err := logrus.ParseLevel(cfg.logLevel)
			if err != nil {
				fmt.Println(err.Error())
			}
			logrus.SetLevel(lvl)
		case "db.connection":
			cfg.dbURI = p.Value
		}
	}
}

type User struct {
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("retrieving request from %s", r.Host)
	logrus.Debug("read request body")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user User
	logrus.Debug("unmarshal request body")
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.Infof("establishing connection wtih %s", cfg.dbURI)
	logrus.Infof("user %s was created successfully", user.Name)
	w.Write([]byte(fmt.Sprintf("user:%s", user.Name)))
}