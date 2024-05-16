package main

import (
	"net/http"

	"github.com/RedrikShuhartRed/go-mongodb-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/user", uc.CreatUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}
