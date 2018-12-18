package graphql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"mollydb/database"
	"mollydb/model"
	"os"
)

var mollyDBMutation *graphql.Object

const (
	uri     = "uri"
	argVerb = "verb"
)

func defineMutation() {

	registerStorage := &graphql.Field{
		Type: storageType,
		Description: "The purpose of this mutation is register new storages" +
			" into mollyDB",
		Args: graphql.FieldConfigArgument{
			storageName: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the storage",
			},
			storagePath: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The path of the storage",
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name, _ := p.Args[storageName].(string)
			path, _ := p.Args[storagePath].(string)
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

	unRegisterStorage := &graphql.Field{
		Type: graphql.String,
		Description: "The purpose of this mutation is unregister an existing" +
			" storage from mollyDB",
		Args: graphql.FieldConfigArgument{
			storageName: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the storage",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			name, _ := p.Args[storageName].(string)
			delete(database.GetInstance().GetStorageDict(), name)
			return "storage deleted successfully!", nil
		},
	}


	argsRestHook:= graphql.FieldConfigArgument{
		uri: &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The uri of the hook",
		},
		argVerb: &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The argVerb of the hook",
		},
		propertyPath: &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The path of the hooked property",
		},
	}


		addRestHook := &graphql.Field{
		Type: graphql.String,
		Description: "The purpose of this mutation is hook property and be" +
			" notified when these have changed",
		Args: argsRestHook,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			uri := argsToString(uri, p)
			verb := argsToString(argVerb, p)
			path := argsToString(propertyPath, p)
			database.GetInstance().AddPropertyRestHook(path, uri, verb)
			return "property hooked", nil
		},
	}

	deleteRestHook := &graphql.Field{
		Type: graphql.String,
		Description: "The purpose of this mutation is deleting an existing" +
			" hook",
		Args: argsRestHook,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			uri := argsToString(uri, p)
			verb := argsToString(argVerb, p)
			path := argsToString(propertyPath, p)
			database.GetInstance().DeletePropertyRestHook(path, uri, verb)
			return "property hooked was deleted!", nil
		},
	}

	mollyDBMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Description: "This permits interact with mollyDB system in order to" +
			" create a new storage",
		Fields: graphql.Fields{
			"register":       registerStorage,
			"unRegister":     unRegisterStorage,
			"addRestHook":    addRestHook,
			"deleteRestHook": deleteRestHook,

		},
	})
}
