package controllers

import (
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Ini halaman get users"))
}
