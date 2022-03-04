package main

import (
	"net/http"
	"pertemuan1/controllers"
)

const port string = ":1234"

func main() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.ListenAndServe(port, nil)
}
