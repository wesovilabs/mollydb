package database

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"mollydb/model"
	"mollydb/util"
	"sync"
)

const (
	connectionPath = "mollydb://%s/%s?key=%s"
)

var instance *Database

var once sync.Once

//GetInstance ensure there's only an instance
func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{
			storageDict:   make(map[string]*model.Storage),
			propertyCache: make(map[string]*model.Property),
			propertyHooks: make(map[string][]Hook),
		}
	})
	return instance
}

//Database structure
type Database struct {
	sync.RWMutex
	storageDict   map[string]*model.Storage
	propertyCache map[string]*model.Property
	propertyHooks map[string][]Hook
}

//AddToStorageDict function to add an element into the storage
func (db *Database) AddToStorageDict(key string, value *model.Storage) {
	db.Lock()
	db.storageDict[key] = value
	db.Unlock()
}

//AddToPropertyCache function to add a property into the cache
func (db *Database) AddToPropertyCache(key string, value *model.Property) {
	db.Lock()
	db.propertyCache[key] = value
	db.Unlock()
}

//AddToPropertyHooks function to create a new hook
func (db *Database) AddToPropertyHooks(key string, value []Hook) {
	db.Lock()
	db.propertyHooks[key] = value
	db.Unlock()
}

//GetStorageDict returns the storage dictionary
func (db *Database) GetStorageDict() map[string]*model.Storage {
	return db.storageDict
}

//GetPropertyCache returns the cache for the properties
func (db *Database) GetPropertyCache() map[string]*model.Property {
	return db.propertyCache
}

//GetPropertyHooks returns the hooks
func (db *Database) GetPropertyHooks() map[string][]Hook {
	return db.propertyHooks
}

//SetUpStorage initialize the storage
func (db *Database) SetUpStorage(s *model.Storage) {
	db.AddToStorageDict(s.Name, s)
	fileList := util.GetFiles(s.Path)
	for _, file := range fileList {
		if dir, _ := util.IsDir(file); !dir {
			db.Add(s, file)
		}
	}
	go WatchPath(s.Path, s)
}

//AddPropertyRestHook creates a new hook
func (db *Database) AddPropertyRestHook(path string, uri string, verb string) {
	log.Printf("Registering hook on %s ", path)
	hook := &RestHook{
		URI:  uri,
		Verb: verb,
	}
	val, ok := db.GetPropertyHooks()[path]
	if !ok {
		log.Printf("It's a new hook")
		db.AddToPropertyHooks(path, []Hook{hook})
		return
	}
	db.AddToPropertyHooks(path, append(val, hook))

}

//DeletePropertyRestHook creates a new hook
func (db *Database) DeletePropertyRestHook(path string, uri string,
	verb string) {
	log.Printf("Registering hook on %s ", path)
	hooks, ok := db.GetPropertyHooks()[path]
	newHooks := []Hook{}
	if ok {
		for _, h := range hooks {
			if h.Type() == restHook {
				hRestHook := h.(RestHook)
				if !(hRestHook.URI == uri && hRestHook.Verb == verb) {
					newHooks = append(newHooks, h)
				}
			}
		}
	}
	db.AddToPropertyHooks(path, newHooks)

}

func load(source string, res interface{}) {
	content, err := ioutil.ReadFile(source)
	if err == nil && util.IsYaml(source) {
		err = yaml.Unmarshal(content, res)
		if err != nil {
			log.Printf("\n%#v\n", err)
			log.Printf("Fails silently %s ", err.Error())
		}
	}
}

//Add function to add a new storage
func (db *Database) Add(storage *model.Storage, source string) {
	data := make(map[string]interface{})
	load(source, &data)
	log.Printf("   [%s] document path: %s", storage.Name, source)
	name, ext, err := util.GetFileName(storage.Path, source)
	defer db.Update(storage, name, ext, data)
	if err != nil {
		log.Printf("\n%#v\n", err)
		log.Printf("It fails silently %s", err.Error())
		return
	}
}

//Delete function to delete a document from a storage
func (db *Database) Delete(storage *model.Storage, source string) {
	name, ext, err := util.GetFileName(storage.Path, source)
	defer db.Update(storage, name, ext, nil)
	if err != nil {
		log.Printf("\n%#v\n", err)
		log.Printf("It fails silently %s", err.Error())
		return
	}
}

//Update function to update a storage
func (db *Database) Update(storage *model.Storage, name string, ext string,
	data map[string]interface{}) {
	db.GetStorageDict()[storage.Name].AddDocument(name, data)
	for _, s := range db.GetStorageDict() {
		for _, d := range s.Documents {
			for _, p := range d.Properties {
				db.updateProperty(s.Name, d.Name, p)
			}
		}
	}

}

func (db *Database) updateProperty(storage string, document string, property *model.Property) {
	path := fmt.Sprintf(connectionPath, storage, document, property)
	v, ok := db.GetPropertyCache()[path]
	if !ok || v != property {
		log.Printf("Property %s is new or has changed its"+
			" value"+
			" hook on", property.Key)
		notify(path)
	}
	db.AddToPropertyCache(path, property)
}
