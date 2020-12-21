package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tommaso-borgato/go-server/util"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, util.GetMessage())
	fmt.Fprintf(w, "\n")
}

func main() {
	fmt.Printf("START\n")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("END\n")
}
