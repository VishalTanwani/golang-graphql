package main

import (
	"fmt"
	"graphql/app"
	"graphql/app/db"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello world")
	os.Setenv("JWI_KEY", "thisIsGraphQLFinalAssignmentafterthaticandocomptetivecoding")
	db.Config()
	app.Start()
	// gql.GetSchema()
}
