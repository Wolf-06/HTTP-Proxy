package main

import (
	intr "HTTPproxy/res"
	"net/http"
)

func main() {
	http.HandleFunc("/", intr.HandleRequest)
	http.HandleFunc("/testcat_an.html", intr.HandleRequest)
	server := intr.Server{":8080"}
	server.Start()
}
