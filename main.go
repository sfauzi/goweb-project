package main

import (
	"goweb-project/config"
	"goweb-project/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// hOMEPAGE
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8081", nil)
}
