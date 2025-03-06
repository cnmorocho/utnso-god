package controllers

import (
	"cnmorocho/utnso-god/src/kernel/services"
	"net/http"
)

type KernelController struct {
	KernelService *services.KernelService
}

func (kc *KernelController) CreateProcess(res http.ResponseWriter, req *http.Request) {
}
