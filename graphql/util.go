package graphql

import "github.com/graphql-go/graphql"

func argsToString(name string, p graphql.ResolveParams) string {
	v,_:= p.Args[name].(string)
	return v
}
