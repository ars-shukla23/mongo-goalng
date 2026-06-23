package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ars-shukla23/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	u := models.User{}
	coll := uc.client.Database("mongo-golang").Collection("users")
	err = coll.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj) // or w.Write(uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}
	u.Id = primitive.NewObjectID()
	coll := uc.client.Database("mongo-golang").Collection("users")
	_, err = coll.InsertOne(context.Background(), u)
	if err != nil {
		fmt.Println(err)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Id := p.ByName("id")
	oid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	coll := uc.client.Database("mongo-golang").Collection("users")
	result, err := coll.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil || result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user %s\n", oid)

}
