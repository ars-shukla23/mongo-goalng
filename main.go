package main

import(
   "github.com/julienschmidt/httprouter"
   "gopkg.in/mgo.v2"
   "net/http"
   "github.com/ars-shukla23/mongo-golang/controllers"
)
 
func main(){
  r:=httprouter.New()
  uc:=controllers.NewUserController(getSession())
  r.GET("/user/:id",uc.GetUser)
  r.POST("/user",uc.CreateUser)
  r.DELETE("/user/:id",uc.DeleteUser)
  http.ListenAndServe("localhost:9000",r)
 
} 

func getSession() *mgo.Session{
   session,err:=mgo.Dial("mongodb://localhost:27017")
   if err!=nil{
	  panic(err)
   }
   return session
}