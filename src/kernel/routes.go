package main

import (
	"cnmorocho/utnso-god/src/kernel/controllers"
	"net/http"
)

func SetupRoutes(router *http.ServeMux) {
	healthController := controllers.HealthController{}
	kernelController := controllers.KernelController{}

	router.HandleFunc("GET /kernel/health", healthController.HealthCheck)
	router.HandleFunc("POST /kernel/process/create", kernelController.CreateProcess)
}
