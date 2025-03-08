package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func StartServer() {
	router := http.NewServeMux()
	port := ":8080"

	SetupRoutes(router)
	slog.Info(fmt.Sprintf("Servidor de kernel iniciado y escuchando en el puerto %s", port))
	http.ListenAndServe(port, router)
}
