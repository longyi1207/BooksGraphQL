package main

import (
	"log"
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var authorType *graphql.Object
var bookType *graphql.Object

func main() {
	bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"name":     &graphql.Field{Type: graphql.String},
			"authorId": &graphql.Field{Type: graphql.Int},
		},
	})

	authorType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					author, ok := params.Source.(Author)
					if ok {
						var authorBooks []Book
						for _, book := range booksData {
							if book.AuthorID == author.ID {
								authorBooks = append(authorBooks, book)
							}
						}
						return authorBooks, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return booksData, nil
				},
			},
			"bookById": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, isOK := params.Args["id"].(int)
					if isOK {
						for _, book := range booksData {
							if int(book.ID) == id {
								return book, nil
							}
						}
					}
					return nil, nil
				},
			},
			"bookByName": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, isOK := params.Args["name"].(string)
					if isOK {
						for _, book := range booksData {
							if book.Name == name {
								return book, nil
							}
						}
					}
					return nil, nil
				},
			},
			"authors": &graphql.Field{
				Type: graphql.NewList(authorType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return authorsData, nil
				},
			},
			"authorById": &graphql.Field{
				Type: authorType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, isOK := params.Args["id"].(int)
					if isOK {
						for _, author := range authorsData {
							if int(author.ID) == id {
								return author, nil
							}
						}
					}
					return nil, nil
				},
			},
			"authorByName": &graphql.Field{
				Type: authorType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, isOK := params.Args["name"].(string)
					if isOK {
						for _, author := range authorsData {
							if author.Name == name {
								return author, nil
							}
						}
					}
					return nil, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	log.Printf("Listening on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
