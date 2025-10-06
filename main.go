package main

import (
	"goweb-project/config"
	"goweb-project/controllers/categorycontroller"
	"goweb-project/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// hOMEPAGE
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8081", nil)
}
