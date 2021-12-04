package main

import (
	"github.com/g1eng/httpfilter/synthesis"
	"log"
	"net/http"
)

func serve() {
	s := http.Server{
		Addr: "127.0.0.1:8080",
	}
	AND := synthesis.AuthAND
	OR := synthesis.AuthOR
	forbid := synthesis.Forbid
	noAuth := synthesis.NoAuth

	http.HandleFunc("/o/ha", OR(forbid, noAuth)(oHandler))
	http.HandleFunc("/o/con", AND(forbid, noAuth)(conchHandler))

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	serve()
}
