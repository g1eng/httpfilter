package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func conchRoute(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	if ps.ByName("ok") != "" {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func oRoute(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	if ps.ByName("ok") != "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
