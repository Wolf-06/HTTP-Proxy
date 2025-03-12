package res

import (
	"fmt"
	"net/http"
)

// Structures
type Server struct {
	Port string
}

//Functions

func (s *Server) Start() {
	addr := s.Port
	fmt.Println("Starting server")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
