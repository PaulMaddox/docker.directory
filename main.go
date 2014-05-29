package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/PaulMaddox/docker.directory/api"
	"labix.org/v2/mgo"
)

func main() {

	// Define CLI arguments
	log.Print("Processing CLI arguments...")
	var (
		port = flag.Int("port", 4000, "HTTP port to listen on")
		db   = flag.String("db", "localhost", "MongoDB database host")
	)

	// Parse CLI arguments
	flag.Parse()

	// Connect to the database
	log.Print("Connecting to database...")
	database, err := mgo.Dial(*db)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB database %s (%s)", *db, err)
	}
	defer database.Close()

	// Start the server
	log.Print("Starting HTTP server...")
	router := api.NewRouter(database)
	bind := "0.0.0.0:" + strconv.Itoa(*port)
	log.Printf("Listening on %s", bind)

	log.Fatal(http.ListenAndServe(bind, router))
}
