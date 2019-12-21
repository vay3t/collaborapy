package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		//fmt.Fprint(w, string(requestDump))
		fmt.Println(string(requestDump))
	}
	fmt.Println("-------------------------------------")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DumpRequest).Methods("GET")
	router.HandleFunc("/", DumpRequest).Methods("POST")
	fmt.Println("Server started on 127.0.0.1:7070")
	fmt.Println("-------------------------------------")
	log.Fatal(http.ListenAndServe(":7070", router))
}
