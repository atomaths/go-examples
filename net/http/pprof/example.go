package main

import "log"
import "net/http"
import _ "net/http/pprof"

func main() {
        log.Println(http.ListenAndServe("localhost:6060", nil))

	// To view all available profiles, open http://localhost:6060/debug/pprof/ in your browser.
}
