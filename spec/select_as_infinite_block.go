package main

import (
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                log.Print("Serving request...")
                // Try not spinning on the cpu, use select {}
                // for {}
                select {}
        })
        log.Fatal(http.ListenAndServe(":8080", nil))
}
