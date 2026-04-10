package main

import (
	"log"
	"net/http"
	"os"

	"LaFinale/pkg/api"
	"LaFinale/pkg/db"

	_ "modernc.org/sqlite"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	api.Init()

	dir := "web"

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Server started on http://localhost:%s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
