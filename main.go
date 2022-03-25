package main

import (
	"github.com/makinap/osu/config"
	"github.com/makinap/osu/migration"

	//"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/makinap/osu/graph"
	"github.com/makinap/osu/graph/generated"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	/*db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			"127.0.0.1", 5432, "postgres", "postgres", "postgres",
		),
	)
	if err != nil {
		log.Fatalln(err)
	}*/
	migration.MigrateTable()

	db := config.GetDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())

	//c := generated.Config{Resolvers: &graph.Resolver{DB: db}}
	//c.Directives.Auth = directives.Auth

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
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