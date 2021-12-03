package responder

import "net/http"

func Write400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func Write401(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func Write403(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}
