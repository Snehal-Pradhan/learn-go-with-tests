package main

import "net/http"

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	http.ListenAndServe(":5000", server)
}
