package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

// Constructor return karta hain Pointer to APIServer
func NewAPISERVER(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := http.NewServeMux()

	// register services
	log.Println("Listing at ", s.addr)
	http.ListenAndServe(s.addr, router)
	return nil
}
