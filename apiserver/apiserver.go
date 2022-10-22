package apiserver

import (
	"github.com/gorilla/mux"
)

type ApiServer struct {
	Adress string
	router *mux.Router
}

type serverOpt func(*ApiServer)

func New(options ...serverOpt) *ApiServer {
	svr := &ApiServer{}
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func AdressOpt(adress string) func(*ApiServer) {
	return func(s *ApiServer) {
		s.Adress = adress
	}
}

func RouterOpt(router *mux.Router) func(*ApiServer) {
	return func(s *ApiServer) {
		s.router = router
	}
}
