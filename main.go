package main

import (
	"fmt"
	"net/http"

	"github.com/RedrikShuhartRed/go-mongodb-api/controllers"
	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Errorf("%w", err)
	}
	defer s.Close()
	return s
}
func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}
