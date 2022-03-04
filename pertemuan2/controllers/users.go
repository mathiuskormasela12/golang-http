package controllers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:username`
	Password string `json:password`
}

type Users []User

func GetUsers(res http.ResponseWriter, req *http.Request) {
	users := Users{
		User{Username: "mathiuscg", Password: "123"},
		User{Username: "chikaplwd3", Password: "chika123"},
	}

	json.NewEncoder(res).Encode(users)
}
