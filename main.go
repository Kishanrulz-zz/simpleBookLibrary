package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kishanrulz/simpleBookLibrary/controller"
	"net/http"
)

func main() {
	r := httprouter.New()
	api := controllers.NewController("mongodb://localhost")

	r.POST("/user", api.CreateUser)
	r.GET("/user/:id", api.GetUser)
	r.PUT("/user/:id", api.UpdateUser)
	r.DELETE("/user/:id", api.DeleteUser)

	r.POST("/book", api.CreateBook)
	r.GET("/book/:id", api.GetBook)
	r.PUT("/book/:id", api.UpdateBook)
	r.DELETE("/book/:id", api.DeleteBook)
	r.POST("/book/:id/issue", api.IssueBook)

	fmt.Print("listening at port 6070")
	http.ListenAndServe("localhost:6070", r)
}
