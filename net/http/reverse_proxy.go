// net/http/fcgi를 써도 되지만,
// 아래처럼 NginX의 proxy_pass를 쓰고
// Go의 http.ListenAndServe() 로 받을 수 있음.
// nginx.conf
/*
    server {
        listen 8080;
        server_name  static.somcloud.com;

        # 이렇게 proxy 방식도 가능
        location ~ /proxy/ {
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host $host;
            proxy_pass http://127.0.0.1:4001;
        }

        location ~ ^/theme/(.*)/pack/ {
            fastcgi_pass    127.0.0.1:4000;
            include         fastcgi_params;
        }

        location ~ / {
            fastcgi_pass    127.0.0.1:3999;
            include         fastcgi_params;
        }
    }
*/

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// 그냥 / 로 pattern을 받아도 됨.
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/proxy/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":4001", nil))
}
