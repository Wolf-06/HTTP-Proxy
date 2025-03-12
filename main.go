package main

import (
	intr "HTTPproxy/res"
	"net/http"
)

func main() {
	go intr.Logger()
	http.HandleFunc("/", intr.HandleRequest)
	http.HandleFunc("/testcat_an.html", intr.HandleRequest)
	http.HandleFunc("/testcat_ac.html", intr.HandleRequest)
	server := intr.Server{":8080"}
	server.Start()
}
