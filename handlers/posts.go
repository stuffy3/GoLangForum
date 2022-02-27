package handlers

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)


func HandleImage(w  http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/octet-stream" )
fmt.Fprintf(w, "Hello there\n")

		b, err := ioutil.ReadFile("6e8.jpg")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		w.Write(b) 
}