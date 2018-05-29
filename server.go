package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/caesarCipher", func(c echo.Context) error {
		plaintext := c.QueryParam("plaintext")
		shift, error := strconv.Atoi(c.QueryParam("shift"))
		if error != nil {
			fmt.Println("Caesar Cipher", error)
		}
		result := caesarEncrypt(plaintext, shift)
		return c.JSON(http.StatusOK, result)
	})

	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	e.GET("/grizzle", echo.WrapHandler(h))
	e.POST("/grizzle", echo.WrapHandler(h))

	e.Logger.Fatal(e.Start(":8000"))
}

// caesarEncrypt is alphabet agnostic.
// TODO: determine alphabet from input
func caesarEncrypt(plaintext string, shift int) string {
	runes := []rune(plaintext)
	for i := range runes {
		runes[i] += int32(shift)
	}
	return string(runes)
}
