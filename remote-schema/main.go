package main

import (
		"log"
		"fmt"

    "github.com/labstack/echo"
		"github.com/labstack/echo/middleware"
		"github.com/99designs/gqlgen/graphql/playground"
		"github.com/99designs/gqlgen/graphql/handler"
		"gorm.io/gorm"
		"remote-schema/config"
		"remote-schema/plugins"
		"remote-schema/graph/generated"
		"remote-schema/graph/resolver"
)


type ItemData struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func main() {    
		config, err := config.Environ()
		if err != nil {
			log.Fatalf("環境変数の設定に失敗しました: %v\n", err)
		}
		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		db, err := plugins.InitWithDSN(config.Database.URL)
		if err != nil {
			fmt.Println("DBへの接続に失敗しました: " + err.Error())
		}
		fmt.Printf("%v\n", db)
	
	
		initRouting(e, db)
	
		// dsn := "host=localhost user=user password=password dbname=local-golang-twitter port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
		
	
		err = e.Start(":3000")
		if err != nil {
			log.Fatalln(err)
		}
    e.Logger.Fatal(e.Start(":3000"))
}

func initRouting(e *echo.Echo, db *gorm.DB) {
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{DB: db}},
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
}