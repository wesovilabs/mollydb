package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var storageConn *relay.GraphQLConnectionDefinitions

func defineConnections() {
	storageConn = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     storage,
		NodeType: storageType,
		ConnectionFields: map[string]*graphql.Field{
			"name": {
				Name:        "name",
				Type:        graphql.String,
				Description: "Document name",
			},
		},
	})
}
