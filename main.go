package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/api/users" 	   , getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}" , getUser).Methods("GET")
	r.HandleFunc("/api/users"      , createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}" , updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}" , deleteUser).Methods("DELETE")

	r.HandleFunc("/api/articles" 	   , getArticles).Methods("GET")
	r.HandleFunc("/api/articles/{id}" , getArticle).Methods("GET")
	r.HandleFunc("/api/articles"      , createArticle).Methods("POST")
	r.HandleFunc("/api/articles/{id}" , updateArticle).Methods("PUT")
	r.HandleFunc("/api/articles/{id}" , deleteArticle).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":9000" , r))
}

func main(){
	InitialDatabase()
	db.AutoMigrate(&User{} , &Article{})
	initializeRouter()

}