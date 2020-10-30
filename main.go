package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World", nil
			},
		},
	},
})

// // calcluate module for calculating two numbers
// func Calculate(a int, b int) int {
// 	return a + b
// }

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	//create a graphql-go http handler with our defined schema
	// and also set to return a pretty JSON
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	//serve a grapqhl endpoint at /graphql
	http.Handle("/graphql", h)

	//and serve!
	http.ListenAndServe(":5000", nil)
}
