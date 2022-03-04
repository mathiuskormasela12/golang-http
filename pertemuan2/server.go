package main

import (
	"net/http"
	"pertemuan2/controllers"
)

func main() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.ListenAndServe(":3000", nil)
}
