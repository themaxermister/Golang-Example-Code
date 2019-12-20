// http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/
// https://freshman.tech/web-development-with-go/
// https://gowebexamples.com/
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/
// Deploy with Heroku
package main

import (
	"flag"
	L "github.com/themaxermister/go-project/lib"
	"log"
	"net/http"
	"os"
)

func main() {
	// API
	L.ApiKey = flag.String("apikey", "", "Newsapi.org access key") //f81239d642ad491e8580161ffa168635
	flag.Parse()

	if *L.ApiKey == "" {
		log.Fatal("apiKey must be set")
	}

	//Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// File Assets
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Pages
	mux.HandleFunc("/search", L.SearchHandler)
	mux.HandleFunc("/", L.IndexHandler)
	http.ListenAndServe(":"+port, mux)
}
