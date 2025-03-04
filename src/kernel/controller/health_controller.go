package controller

import "net/http"

type HealthController struct{}

func (h *HealthController) HealthCheck(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("OK"))
}
