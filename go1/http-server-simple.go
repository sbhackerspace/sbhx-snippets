package main

import "net/http"

const PORT = "8082"

func main() {
	println("Visit http://localhost:" + PORT)
	http.Handle("/", http.HandlerFunc(handleRequest))
	http.ListenAndServe(":"+PORT, nil)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello from Go server"))
}
