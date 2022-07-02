package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bharathreddy-97/Manga-Dispencer/constants"
	"github.com/bharathreddy-97/Manga-Dispencer/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Initialising Server")
	// Register all the urls and handlers
	router := mux.NewRouter()
	fmt.Println("Registering End points")
	router.HandleFunc(constants.GetAllbooks, handlers.HandleGetAllBooksList).Methods(http.MethodGet)
	router.HandleFunc(constants.GetPageOfBook, handlers.HandleGetPage).Methods(http.MethodGet)

	fmt.Println("Registered all End points")
	fmt.Println("Listenening on port 3232")
	log.Fatal(http.ListenAndServe(":3232", router))
}
