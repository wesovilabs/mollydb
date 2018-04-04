package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"mollydb/database"
	"mollydb/model"
)

const (
	query = "Query"
	any   = "any"
)

var mollyDBQuery *graphql.Object

func defineQueries() {

	fieldStorages := &graphql.Field{
		Name: "QueryStorages",
		Description: "This query allow us to deep from the root of the" +
			" mollyDB system until a Property definition.",
		Type: &graphql.List{OfType: storageType},
		Args: map[string]*graphql.ArgumentConfig{
			"name": {
				Type:         graphql.String,
				DefaultValue: any,
				Description: "The name of the storage. " +
					"StorageList can be filtered by name",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var storageList []*model.Storage
			for _, s := range database.GetInstance().GetStorageDict() {
				storageList = append(storageList, s)
			}
			return storageList, nil
		},
	}

	fieldProperty := &graphql.Field{
		Name: "QueryProperty",
		Description: "Find a property in any document of any storage by the" +
			" connection path. This is an unique value for each property in" +
			" all the mollydb system. The output is a single Property",
		Type: propertyType,
		Args: map[string]*graphql.ArgumentConfig{
			"path": {
				Type:        graphql.String,
				Description: "The path of the property",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			propertyPath := p.Args["path"].(string)
			fmt.Println(database.GetInstance().GetPropertyCache())
			return database.GetInstance().GetPropertyCache()[propertyPath], nil
		},
	}

	fieldProperties := &graphql.Field{
		Type: &graphql.List{OfType: propertyType},
		Name: "QueryProperties",
		Description: "Find properties by filtering records by the name" +
			" of the storage or/and" +
			" the name of the document or/and the key of the property. " +
			"\nDefault value for filters is any. " +
			"The output is an array of type Property",
		Args: map[string]*graphql.ArgumentConfig{
			"storage": {
				Type:         graphql.String,
				DefaultValue: any,
				Description:  "The name of the storage",
			},
			"document": {
				Type:         graphql.String,
				DefaultValue: any,
				Description:  "The name of the document",
			},
			"property": {
				Type:         graphql.String,
				DefaultValue: any,
				Description:  "The key of the property",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			storage := p.Args["storage"].(string)
			document := p.Args["document"].(string)
			property := p.Args["property"].(string)
			var result []*model.Property
			for _, s := range database.GetInstance().GetStorageDict() {
				if storage == any || storage == s.Name {
					for _, d := range s.Documents {
						if document == any || document == d.Name {
							for _, k := range d.Properties {
								if property == any || property == k.Key {
									result = append(result, k)
								}
							}
						}
					}
				}
			}
			return result, nil
		},
	}

	mollyDBQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: query,
		Fields: graphql.Fields{
			"storageList": fieldStorages,
			"properties":  fieldProperties,
			"property":    fieldProperty,
		},
	})
}
