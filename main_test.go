package main

/**
1. Module is for testing server
*/

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func TestGrapqlServer(t *testing.T) {
	// query
	query := "{ latestPost }"
	//expected output
	expected := map[string]interface{}{
		"latestPost": "Hey World",
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
	})

	// if we encounter an error
	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors %v", result.Errors)
	}

	if !reflect.DeepEqual(result.Data, expected) {
		t.Fatalf("wrong result, query: %v graphql result %v", query, testutil.Diff(expected, result))
	}
}
