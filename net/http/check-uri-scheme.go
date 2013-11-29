// https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
// https://plus.google.com/117756050374403495170/posts/7uoVL2fdRPS

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var scheme string
		if r.TLS == nil {
			scheme = "HTTP"
		} else {
			scheme = "HTTPS"
		}
		fmt.Fprintf(w, "Your request's scheme is "+scheme)
	})

	log.Fatal(http.ListenAndServe(":9999", nil))
}
