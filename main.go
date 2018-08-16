package main

import (
	context "context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"

	"github.com/julienbreux/gqlgentest/schema/graph"
	"github.com/julienbreux/gqlgentest/schema/model"
)

func main() {
	http.Handle("/", handler.Playground("Test", "/query"))
	http.Handle("/query", handler.GraphQL(
		graph.NewExecutableSchema(new()),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func new() graph.Config {
	return graph.Config{
		Resolvers: &resolver{},
	}
}

type resolver struct{}

func (r *resolver) Query() graph.QueryResolver {
	return &queryResolver{}
}

type queryResolver struct{}

func (qr *queryResolver) Get(ctx context.Context, name *string) (*model.ObjectReturned, error) {
	return &model.ObjectReturned{
		Name: name,
	}, nil
}
