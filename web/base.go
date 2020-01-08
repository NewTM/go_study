package main

import (
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	str := "hello world------"
	w.Write([]byte(str))
	// fmt.Fprintf(w, "hello world")
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/test", test)
	server.ListenAndServe()
}
