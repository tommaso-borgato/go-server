package main

import (
	"fmt"
	"net/http"

	"github.com/tommaso-borgato/go-server/util"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, util.GetMessage())
	fmt.Fprintf(w, "\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
