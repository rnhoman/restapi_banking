package main

import (
	"net/http"
)

func Start()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	// log.Fatal(http.ListenAndServe("localhost:8888", mux))
	http.ListenAndServe("localhost:8888", mux)
}