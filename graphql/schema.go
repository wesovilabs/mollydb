package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

//MollySchema graphql schema
var MollySchema graphql.Schema

func init() {
	defineTypes()
	//defineConnections()
	defineQueries()
	defineMutation()

	var err error

	MollySchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    mollyDBQuery,
		Mutation: mollyDBMutation,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
