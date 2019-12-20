package main

import (
	R "github.com/themaxermister/another-go/routes"
	"net/http"
	"time"
)

func main() {
	R.p("ChitChat", R.Version(), "started at", R.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(R.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", R.Index)
	// error
	mux.HandleFunc("/err", R.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", R.Login)
	mux.HandleFunc("/logout", R.Logout)
	mux.HandleFunc("/signup", R.Signup)
	mux.HandleFunc("/signup_account", R.SignupAccount)
	mux.HandleFunc("/authenticate", R.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", R.NewThread)
	mux.HandleFunc("/thread/create", R.CreateThread)
	mux.HandleFunc("/thread/post", R.PostThread)
	mux.HandleFunc("/thread/read", R.ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:           R.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(R.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(R.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
