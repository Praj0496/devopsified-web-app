package main

import (
    "net/http"
    "log"
    "os"
)

func main() {
    // Serve all static files (CSS, JS, images, index.html) from the "static" directory
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
