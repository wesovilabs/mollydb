package api

import (
	"github.com/graphql-go/handler"
	"mollydb/graphql"
)

const path = "/graphql"

//Handler function to define graphql endpoint
func Handler() (string, *handler.Handler) {
	h := handler.New(&handler.Config{
		Schema:   &graphql.MollySchema,
		Pretty:   true,
		GraphiQL: true,
	})
	return path, h
}
