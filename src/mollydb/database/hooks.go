package database

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	httpJSON          = "application/json"
	httpHeaderContent = "Content-type"
	restHook          = "restHook"
)

//Hook interface
type Hook interface {
	//Launch function to be implemented
	Launch(path string)
	//Type returns the hook type
	Type() string
}

//RestHook structure
type RestHook struct {
	//URI attribute
	URI string `json:"uri"`
	//Verb attribute
	Verb string `json:"verb"`
}

//Launch function omplementation for RestHooks
func (wh RestHook) Launch(path string) {
	log.Printf("notifying rest hook on %s %s that property %s has changed ",
		wh.Verb, wh.URI, path)
	client := http.Client{}
	property := GetInstance().GetPropertyCache()[path]
	b, err := json.Marshal(property)
	property.Update()
	if err != nil {
		log.Fatalf("Unexpected error %s", err.Error())
		return
	}
	req, err := http.NewRequest(wh.Verb, wh.URI, bytes.NewBuffer(b))
	req.Header.Set(httpHeaderContent, httpJSON)
	client.Do(req)
}

//Type function
func (wh RestHook) Type() string {
	return restHook
}
func notify(path string) {
	hooks, ok := GetInstance().GetPropertyHooks()[path]
	if ok {
		log.Printf("Notifying hooks on %s", path)
		for _, h := range hooks {
			go h.Launch(path)
		}
	}
}
