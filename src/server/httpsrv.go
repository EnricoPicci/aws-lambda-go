package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/serverless/helloYou/src/core"
	"github.com/serverless/helloYou/src/events"
)

// https://stackoverflow.com/a/56312831/5699993
var addr = flag.String("addr", ":8080", "http service address")

// Start the server
func StartHttpSrv() {
	fmt.Printf("Server started on port %v\n", *addr)
	flag.Parse()

	http.HandleFunc("/", homePage)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w)
		if r.Method == "OPTIONS" {
			return
		}
		decoder := json.NewDecoder(r.Body)
		var event events.Hello
		err := decoder.Decode(&event)
		if err != nil {
			log.Fatalf("Error while reading body of hello request %v", err)
		}
		fmt.Printf("Hello received for %v\n", event.Name)
		resp := core.Hello(event.Name)

		fmt.Fprint(w, resp)
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		setupHeader(w)
		if r.Method == "OPTIONS" {
			return
		}
		decoder := json.NewDecoder(r.Body)
		var event events.World
		err := decoder.Decode(&event)
		if err != nil {
			log.Fatalf("Error while reading body of world request %v", err)
		}
		fmt.Printf("World received for %v\n", event.Name)
		resp := core.World(event.Name)

		fmt.Fprint(w, resp)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func setupHeader(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
