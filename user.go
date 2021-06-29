package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


type User struct{
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Article []Article `json:"article"`
}


//db.AutoMigrate(&User{})

func getUsers(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func getUser(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user , params["id"])
	db.Model(&user).Association("Article").Find(&user.Article) //find(xxx) >>> use this struct and put it right here
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user) //get body and struct in this user 
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user , params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var user User
	db.Delete(&user , params["id"])
	json.NewEncoder(w).Encode("user deleted successfully")
}