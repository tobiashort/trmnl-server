package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DisplayResponse struct {
	Status         int    `json:"status"`
	ImageURL       string `json:"image_url"`
	FileName       string `json:"filename"`
	UpdateFirmware bool   `json:"update_firmware"`
	FirmwareURL    string `json:"firmware_url"`
	RefreshRate    string `json:"refresh_rate"`
	ResetFirmware  bool   `json:"reset_firmware"`
}

func NewDisplayResponse() DisplayResponse {
	return DisplayResponse{
		Status:         0,
		ImageURL:       "http://192.168.1.122:8080/image",
		FileName:       fmt.Sprintf("%d", time.Now().Unix()),
		UpdateFirmware: false,
		FirmwareURL:    "",
		RefreshRate:    "10",
		ResetFirmware:  false,
	}
}

type DisplayHandler struct{}

func (DisplayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewDisplayResponse())
}
