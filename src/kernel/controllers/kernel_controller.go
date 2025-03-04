package controllers

import (
	"cnmorocho/utnso-god/src/kernel/entities"
	"net/http"
)

type KernelController struct {
	Kernel *entities.Kernel
}

func (kc *KernelController) CreateProcess(res http.ResponseWriter, req *http.Request) {
	kc.Kernel.CreateProcess()
}
