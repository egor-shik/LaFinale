package main 

import (
    "os"
    "net/http"
    "log"

    "LaFinale/pkg/db"
    "LaFinale/pkg/api"
    _ "modernc.org/sqlite"
)

func main() {
    err := db.Init("scheduler.db")
if err != nil {
	log.Fatal(err)
}

    api.Init()

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