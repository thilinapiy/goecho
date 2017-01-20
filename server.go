package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Version string
	Msg     string
	Time    string
}

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/echo", Index)
	router.HandleFunc("/echo/{msg}", Echo)
	log.Printf("Starting server on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "usage : /echo/{message}")
}

func Echo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := vars["msg"]

	response := Response{
		Version: "v1",
		Msg:     msg,
		Time:    time.Now().Format(time.RFC3339),
	}

	body, _ := json.Marshal(response)
	fmt.Fprintln(w, string(body))
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
