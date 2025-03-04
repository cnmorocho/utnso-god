package main

import "net/http"

func StartServer() {
	router := http.NewServeMux()

	http.Handle("/kernel", router)
	SetupRoutes(router)
	
	http.ListenAndServe(":8080", router)
}
