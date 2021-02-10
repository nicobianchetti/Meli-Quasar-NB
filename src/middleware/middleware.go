package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

//FunctionLog . function tipo handler
type FunctionLog func(http.ResponseWriter, *http.Request)

//Log .
func Log(f FunctionLog) FunctionLog {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Petición: %q, método: %q, headers %q, query parms %q", r.URL.Path, r.Method, r.Header, r.URL.RawQuery)
		f(w, r)
	}
}

// Authentication .
func Authentication(f FunctionLog) FunctionLog {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("api-key")
		if token != "un-token-muy-seguro" {
			//responder "ingreso no autorizado"
			forbidden(w, r)
			return
		}

		f(w, r)
	}
}

//LogAndAuthentication .
func LogAndAuthentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return Log(Authentication(f))
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode("api-key no válida.")
}
