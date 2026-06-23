package main

import (
	"context"
	"net/http"

	"github.com/ars-shukla23/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	client := getClient()
	defer client.Disconnect(context.Background())

	uc := controllers.NewUserController(client)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}
