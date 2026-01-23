package main

import (
	"net/http"

	"gee/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, Gee!"))
	})
	http.ListenAndServe(":8080", r)
}
