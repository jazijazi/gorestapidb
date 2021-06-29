package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Article struct{
	gorm.Model
	Title  string  `json:"title"`
	Content   string  `json:"content"`
	UserID uint `json:"userid"`
}

func getArticles(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var article []Article
	db.Find(&article)
	json.NewEncoder(w).Encode(&article)

}
func getArticle(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var article Article
	db.First(&article , params["id"])
	json.NewEncoder(w).Encode(&article)
}
func createArticle(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	db.Create(&article)
	json.NewEncoder(w).Encode("fffffff")
}
func updateArticle(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var article Article
	db.First(&article , params["id"])
	json.NewDecoder(r.Body).Decode(&article)
	db.Save(&article)
	json.NewEncoder(w).Encode(article)
}
func deleteArticle(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	var article Article
	db.Delete(&article , params["id"])
	json.NewEncoder(w).Encode("article deleted successfully")
}