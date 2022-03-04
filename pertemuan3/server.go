package main

import (
	"net/http"
	"pertemuan3/controllers"
	"pertemuan3/middlewares"
)

func main() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/post/users", middlewares.CheckBody(controllers.CreateUser))
	http.ListenAndServe(":3000", nil)
}