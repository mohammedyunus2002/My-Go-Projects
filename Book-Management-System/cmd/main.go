package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedyunus2002/demo-repo/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println(" Listning on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
