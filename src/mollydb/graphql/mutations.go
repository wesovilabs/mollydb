package graphql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"mollydb/database"
	"mollydb/model"
	"os"
)

var mollyDBMutation *graphql.Object

func defineMutation() {

	fieldRegisterStorage := &graphql.Field{
		Type: storageType,
		Description: "The purpose of this mutation is register new storages" +
			" into mollyDB",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the storage",
			},
			"path": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The path of the storage",
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name, _ := p.Args["name"].(string)
			path, _ := p.Args["path"].(string)
			if _, err := os.Stat(path); err == nil {
				storage := &model.Storage{
					Name:      name,
					Path:      path,
					Documents: make(map[string]*model.Document),
				}
				database.GetInstance().SetUpStorage(storage)
				return storage, nil
			}
			return nil, errors.New("Invalid path")
		},
	}

	fieldUnRegisterStorage := &graphql.Field{
		Type: graphql.String,
		Description: "The purpose of this mutation is unregister an existing" +
			" storage from mollyDB",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the storage",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name, _ := p.Args["name"].(string)
			delete(database.GetInstance().GetStorageDict(), name)
			return "storage deleted successfully!", nil
		},
	}

	fieldPropertyRestHook := &graphql.Field{
		Type: graphql.String,
		Description: "The purpose of this mutation is hook property and be" +
			" notified when these have changed",
		Args: graphql.FieldConfigArgument{
			"uri": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The uri of the hook",
			},
			"verb": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The verb of the hook",
			},
			"path": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The path of the hooked property",
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			uri, _ := p.Args["uri"].(string)
			verb, _ := p.Args["verb"].(string)
			path, _ := p.Args["path"].(string)
			database.GetInstance().AddPropertyRestHook(path, uri, verb)
			return "property hooked", nil
		},
	}

	mollyDBMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Description: "This permits interact with mollyDB system in order to" +
			" create a new storage",
		Fields: graphql.Fields{
			"register":         fieldRegisterStorage,
			"unRegister":       fieldUnRegisterStorage,
			"propertyRestHook": fieldPropertyRestHook,
		},
	})
}
