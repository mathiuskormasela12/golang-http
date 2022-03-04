package controllers

import (
	"net/http"
	"encoding/json"
	// "fmt"
	// "io/ioutil"
)

type User struct {
	Username string `json:username`
	Password string `json:password`
}

type Users []User

var userLists = Users{
	User{ Username: "mathiuscg", Password: "1234" },
	User{ Username: "chikaPLWD3", Password: "chika123" },
}

func GetUsers(res http.ResponseWriter, req *http.Request) {

	// untuk mengentikan sebuah kode dan mencetak sesuatu
	// panic(req.Method)
	if req.Method == "GET" {
		json.NewEncoder(res).Encode(userLists)
	} else {
		// untuk menghentikan program dan menampilkan error
		http.Error(res, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Invalid Request Method", http.StatusMethodNotAllowed)
	} else {
		// menampilkan req.Body
		// var body, err = ioutil.ReadAll(req.Body)
		
		// if err != nil {
		// 	http.Error(res, "Can't read the body request", http.StatusInternalServerError)
		// } else {
		// 	// konver body nya jadi string dari json
		// 	res.Write([]byte(string(body)))
		// }
		// var newUser User

		// mengubah req.Body menjadi string dari json dan assign ke variable newUser
		// err := json.NewDecoder(req.Body).Decode(&newUser)

		// if err != nil {
		// 	fmt.Println("Can't read the body request")
		// }

		// json.NewEncoder(res).Encode(newUser)

		var newUser User
		err := json.NewDecoder(req.Body).Decode(&newUser)

		if err != nil {
			http.Error(res, "Can't read the request body", http.StatusInternalServerError)
			panic(err)
		} else {
			userLists = append(userLists, newUser)

			res.Write([]byte("Berhasil menambah data"))
		}
	}
}