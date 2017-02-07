package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside http handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func startWebServer() {
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
