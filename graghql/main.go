package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Define GraphQL schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, World!", nil
				},
			},
		},
	}),
})

// Define GraphQL handler
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		log.Printf("Errors: %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", graphqlHandler)

	fmt.Println("Server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
