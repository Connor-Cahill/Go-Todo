package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/mongodb/mongo-go-driver/mongo"
)

//sets up echo server
func setupServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yo Yo Yo")
	})
	e.Logger.Fatal(e.Start(":3000"))
}

//TodoItem is the object for Todo List Items
type TodoItem struct {
	name        string
	description string
	importance  int
}

//Sets up mongodb instance
func setupDB() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo!")

	// collection := client.Database("go-todo").Collection("TodoItems")
}

//SetupServer starts the server on localhost:3000
func main() {
	setupServer()
	setupDB()

}
