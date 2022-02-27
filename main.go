package main

import (
	"forum/routes"

	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server started")
	r := routes.Routes{}
	r.Init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
	
	


	