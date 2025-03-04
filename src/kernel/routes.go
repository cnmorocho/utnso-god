package main

import (
	"cnmorocho/utnso-god/src/kernel/controller"
	"net/http"
)

func SetupRoutes(router *http.ServeMux) {
	healthController := controller.HealthController{}
	kernelController := controller.KernelController{}

	router.HandleFunc("GET /health", healthController.HealthCheck)
	router.HandleFunc("POST /process/create", kernelController.CreateProcess)
}
