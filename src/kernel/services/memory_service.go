package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MemoryService struct{}

func NewMemoryService() *MemoryService {
	return &MemoryService{}
}

type RequestCreateProcessDTO struct {
	Size int `json:"size"`
}

type ResponseCreateProcessDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (ms *MemoryService) CreateProcess(req RequestCreateProcessDTO) (ResponseCreateProcessDTO, error) {
	url := ""
	bodyReq, err := json.Marshal(req)
	if err != nil {
		return ResponseCreateProcessDTO{}, fmt.Errorf("error al intentar convertir request a json: %v", err)
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(bodyReq))

	if err != nil {
		return ResponseCreateProcessDTO{}, fmt.Errorf("error al intentar crear proceso en memoria: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseCreateProcessDTO{}, fmt.Errorf("error al intentar crear proceso en memoria: %v", err)
	}

	responseData := ResponseCreateProcessDTO{}
	json.Unmarshal(body, &responseData)

	return responseData, nil
}
