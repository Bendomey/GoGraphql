package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Fetch Posts",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type:        graphql.String,
			Description: "Fetch Latest Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				return "Hey World", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

// // calcluate module for calculating two numbers
// func Calculate(a int, b int) int {
// 	return a + b
// }

func main() {

	//create a graphql-go http handler with our defined schema
	// and also set to return a pretty JSON
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	//serve a grapqhl endpoint at /graphql
	http.Handle("/graphql", h)

	//and serve!
	port := ":5000"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	errServer := http.ListenAndServe(port, nil)
	if errServer != nil {
		log.Fatalf("Error occured while serving graphql server, %v", errServer)
	}
}
