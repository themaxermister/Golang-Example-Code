package main

import (
	"github.com/gorilla/mux"
	L "github.com/themaxermister/multi-page/lib"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", L.Home)
	r.HandleFunc("/contact", L.Contact)
	r.HandleFunc("/faq", L.Faq)

	// Error Handling
	var h http.Handler = http.HandlerFunc(L.ErrorNotFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
