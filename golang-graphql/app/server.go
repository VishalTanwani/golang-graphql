package app

import (
	"context"
	"encoding/json"
	"fmt"
	"graphql/app/gql"
	"net/http"

	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is a graphql project")
}

//ExecuteQuery is the function
func ExecuteQuery(ctx context.Context, query string, variables map[string]interface{}) (*graphql.Result, error) {
	schema, err := gql.GetSchema()
	if err != nil {
		return nil, err
	}

	result := graphql.Do(graphql.Params{
		Schema:         *schema,
		RequestString:  query,
		Context:        ctx,
		VariableValues: variables,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("unxpected errors inside query : %v\n", result.Errors)
	}

	return result, nil

}

func gqlHangler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "no query data", 404)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "error parsing JSON request body", 404)
			return
		}

		result, err := ExecuteQuery(r.Context(), rBody.Query, rBody.Variables)

		if len(result.Errors) > 0 {
			fmt.Printf("unxpected errors inside query : %v\n", result.Errors)
			panic(result.Errors)
		}

		b, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s\n", b)

	})

}

//Start this is the function where it start everything
func Start() {
	fmt.Println("Starting the server")
	http.HandleFunc("/", home)
	http.HandleFunc("/graphql", gqlHangler())
	http.ListenAndServe(":5000", nil)
}
