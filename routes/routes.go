package routes

import (
	"forum/handlers"

	"net/http"
)

type Routes struct {}

func (r Routes) Init() {http.HandleFunc("/", handlers.HandleImage)}
