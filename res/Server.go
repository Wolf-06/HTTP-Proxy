package res

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) Start() {
	addr := s.Port
	fmt.Println("Starting server")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func LogRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved a request: ")
	log.Printf("Method; %s\n", r.Method)
	log.Printf("URL: %s\n", r.URL)
	log.Println("Header: ")
	for key, value := range r.Header {
		log.Printf("%s : %v", key, value)
	}
	fmt.Fprintf(w, "Request Logged\n")
}
