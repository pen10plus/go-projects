package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/pen10plus/Book-store/pkg/routes"
	"github.com/pen10plus/Book-store/pkg/models"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	// Initialize database
	models.InitDB()
	
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
	
}