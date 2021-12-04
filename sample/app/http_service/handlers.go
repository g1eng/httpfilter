package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func conchHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func conchRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if ps.ByName("ok") != "" {
		conchHandler(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func oHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func oRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if ps.ByName("ok") == "moko2" {
		oHandler(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
