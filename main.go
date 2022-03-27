package main

import (
	"github.com/makinap/osu/config"
	"github.com/makinap/osu/directives"
	"github.com/makinap/osu/middlewares"
	"github.com/makinap/osu/migration"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/makinap/osu/graph"
	"github.com/makinap/osu/graph/generated"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {



	migration.MigrateTable()

	db := config.GetDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middlewares.AuthMiddleware)

	e.GET("/", welcome())

	c := generated.Config{Resolvers: &graph.Resolver{DB: db}}
	c.Directives.Auth = directives.Auth

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema( c,
			//generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err := e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}