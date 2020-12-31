package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/gorilla/mux"
)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("\t1.1.1.1\n----------")
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
	host := flag.String("host", "127.0.0.1", "host address")
	port := flag.Int("port", 8000, "host port")

	flag.Parse()

	portString := strconv.Itoa(*port)

	router := mux.NewRouter()
	router.HandleFunc("/", DumpRequest).Methods("GET")
	router.HandleFunc("/", DumpRequest).Methods("POST")
	fmt.Println("Server started on " + *host + ":" + portString)
	fmt.Println("-------------------------------------")
	log.Fatal(http.ListenAndServe(*host+":"+portString, router))
}
