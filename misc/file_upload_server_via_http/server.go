// See http://stackoverflow.com/questions/20765859/go-accepting-http-post-multipart-files
// 
// Server
// 	$ go run server.go
//
// Client
//	$ curl -i -F file=@gopher.png http://127.0.0.1:8080/
//
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	out, err := os.Create("./uploaded.file")
	if err != nil {
		fmt.Fprintf(w, "Failed to open the file for writing")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	// the header contains useful info, like the original file name
	fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}
