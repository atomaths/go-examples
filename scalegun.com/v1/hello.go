package main

import (
	"fmt"
	"net/http"

	"scalegun.com/v1"
)

func first(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, Scalegun")
}

func main() {
	app := scalegun.NewApp("go.scalegun.com", "scalegun_secret")
	app.HandleFunc("/", first)
	app.ListenAndServe()
}
