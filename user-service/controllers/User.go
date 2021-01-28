package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hellorichardpham/onlyfarms/user-service/models"
	"github.com/hellorichardpham/onlyfarms/user-service/models/daos"
)

//User struct
type User struct {
	UserDAO *daos.User
}

//GetUserByID placeholder
func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Error converting id parameter to an integer.")
		return
	}
	users, err := u.UserDAO.Get(id)
	if err != nil {
		fmt.Println("Could not find a user with id:", id, " error:", err)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

//CreateUser placeholder
func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Could decode the payload. Error:", err)
	}
	err = u.UserDAO.Insert(user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("There was a problem inserting the user. Error:", err)
	}
	fmt.Println("Successfully created a new user")
}

//AuthenticateUser placeholder
func (u *User) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Could decode the payload. Error:", err)
	}
	id, err := u.UserDAO.Authenticate(user.Email, user.Password)
	if err != nil {
		fmt.Println("The email or password is invalid. err:", err)
	} else {
		fmt.Println("Successfully authenticated user with id:", id)
	}
}
