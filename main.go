package main

import (
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/api"
)

func main() {

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":4001", router))

}
