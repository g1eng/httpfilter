package main

import (
	"github.com/g1eng/httpfilter/synthesis"
	"github.com/g1eng/httpfilter/synthesis/rt_synthesis"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Serve() *http.Server {
	s := http.Server{
		Addr: "127.0.0.1:8080",
	}
	AND := synthesis.AuthAND
	OR := synthesis.AuthOR
	forbid := synthesis.Forbid
	noAuth := synthesis.NoAuth

	http.HandleFunc("/o/ha", OR(forbid, noAuth)(oHandler))
	http.HandleFunc("/o/con", AND(forbid, noAuth)(conchHandler))

	return &s
}

func ServeRouter() *http.Server {
	AND := rt_synthesis.AuthAND
	OR := rt_synthesis.AuthOR
	unAuth := rt_synthesis.UnAuth
	noAuth := rt_synthesis.NoAuth

	route := httprouter.New()
	route.GET("/o/ha", OR(unAuth, noAuth)(oRoute))
	route.GET("/o/con", AND(unAuth, noAuth)(conchRoute))

	return &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: route,
	}
}
