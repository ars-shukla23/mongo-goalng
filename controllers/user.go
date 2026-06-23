package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"	
	"github.com/ars-shukla23/mongo-golang/models"
)
type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	Id:=p.ByName("id")

	if !bson.IsObjectIdHex(Id){
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid:=bson.ObjectIdHex(Id)
	u:=models.User{}
	err:=uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uj,err:=json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",uj) // or w.Write(uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u:=models.User{}
	err:=json.NewDecoder(r.Body).Decode(&u)
	if err!=nil{
		fmt.Println(err)
	}
	u.Id=bson.NewObjectId()
	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj,err:=json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s\n",uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	Id:=p.ByName("id")
	if !bson.IsObjectIdHex(Id){
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid:=bson.ObjectIdHex(Id)
	err:=uc.session.DB("mongo-golang").C("users").RemoveId(oid)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Deleted user %s\n",oid)
	
	
}
