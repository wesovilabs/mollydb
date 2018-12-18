package model

import (
	"fmt"
	"mollydb/util"
	"strings"
)

//Property structure
type Property struct {
	//Path field
	Path string `json:"path"`
	//Key field
	Key string `json:"key"`
	//Value field
	Value interface{} `json:"value"`
}
//Update content
func (p *Property) Update() {
	//	if p.Value
}

//Properties array of properties
type Properties []*Property

func cleanKeys(elements []string, parent string) []string {
	var result []string
	for _, e := range elements {
		if !(strings.HasPrefix(parent, e) && len(parent) > len(e)) {
			result = append(result, e)
		}
	}
	return result
}

//ExtractKeys function to find the keys
func ExtractKeys(base string, parents []string, value interface{}) []string {
	switch val := value.(type) {
	case map[string]interface{}:
		for k, v := range val {
			path := k
			if len(base) > 0 {
				path = base + "." + path
			}
			parents = append(parents, path)
			parents = ExtractKeys(path, parents, v)
		}
	case map[interface{}]interface{}:
		for k, v := range val {
			path := k.(string)
			if len(base) > 0 {
				path = base + "." + path
			}
			parents = append(parents, path)
			parents = ExtractKeys(path, parents, v)
		}
	}
	for _, p := range parents {
		parents = cleanKeys(parents, p)
	}
	return parents
}

//Document structure
type Document struct {
	//Name field
	Name string `json:"name"`
	//Properties field
	Properties Properties `json:"properties"`
}

//Storage structure
type Storage struct {
	//Name field
	Name string `json:"name"`
	//Path field
	Path string `json:"path"`
	//Documents field
	Documents map[string]*Document `json:"documents"`
}

func path(storage, document, key string) string {
	return fmt.Sprintf("mollydb://%s/%s?key=%s", storage, document, key)
}

//AddDocument adding anew storage
func (f *Storage) AddDocument(documentName string,
	content map[string]interface{}) {
	var parents []string
	var properties Properties
	keys := ExtractKeys("", parents, content)
	for _, k := range keys {
		v := util.Get(content, strings.Split(k, "."))
		propertyPath := path(f.Name,
			documentName, k)
		property := &Property{Path: propertyPath,
			Key:   k,
			Value: v}
		properties = append(properties, property)

	}
	if len(f.Documents) <= 0 {
		f.Documents = make(map[string]*Document)
	}
	//if _, ok := f.Documents[documentName]; !ok {
	f.Documents[documentName] = &Document{
		Name:       documentName,
		Properties: properties,
	}
	//}
}

//DeleteDocument deleting a document from the storage
func (f *Storage) DeleteDocument(documentName string) {
	if _, ok := f.Documents[documentName]; ok {
		f.Documents[documentName] = nil
	}
}
