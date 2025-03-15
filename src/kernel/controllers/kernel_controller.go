package controllers

import (
	"cnmorocho/utnso-god/src/kernel/dtos"
	"cnmorocho/utnso-god/src/kernel/services"
	"encoding/json"
	"log/slog"
	"net/http"
)

type KernelController struct {
	KernelService *services.KernelService
}

func (kc *KernelController) CreateProcess(res http.ResponseWriter, req *http.Request) {
	var createProcessRequest dtos.CreateProcessRequestDTO
	err := json.NewDecoder(req.Body).Decode(&createProcessRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		slog.Error("No se pudo crear el proceso")
		return
	}
	res.Header().Set("Content-Type", "application/json")
	kc.KernelService.CreateProcess(createProcessRequest)
	json.NewEncoder(res).Encode(createProcessRequest)
}
