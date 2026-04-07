package main 

import (
    "os"
    "net/http"
    "log"

    "LaFinale/pkg/db"
)

func main() {
    err := db.Init("scheduler.db")
if err != nil {
	log.Fatal(err)
}
    dir := "web"

    port := os.Getenv("TODO_PORT")
    if port == "" {
        port = "7540"
    }
http.Handle ("/", http.FileServer(http.Dir(dir)))
log.Printf("Server started on http://localhost:%s", port)
    err = http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal(err)
    }

}