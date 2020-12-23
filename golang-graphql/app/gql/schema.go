package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

//GetSchema is will return all query and mutaions
func GetSchema() (*graphql.Schema, error) {
	rootQuery := NewRootQuery()
	rootMutation := NewRootMutation()

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery.Query,
			Mutation: rootMutation.Mutation,
		},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
		return nil, err
	}
	return &sc, nil
}
