package main

import (
	intr "HTTPproxy/res"
	"net/http"
)

func main() {
	http.HandleFunc("/", intr.LogRequest)

	server := intr.Server{":8080"}
	server.Start()
}
