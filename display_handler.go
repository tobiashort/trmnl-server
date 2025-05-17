package main

import (
	"encoding/json"
	"fmt"
	"math"
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

type DisplayHandler struct {
	BaseUrl string
}

func (h DisplayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ts := time.Now().Unix()
	rs := 300 - (ts-150)%300
	res := DisplayResponse{
		Status:         0,
		ImageURL:       h.BaseUrl + "/image",
		FileName:       fmt.Sprintf("%d", int(math.Floor(float64(ts)/300))*300+150),
		UpdateFirmware: false,
		FirmwareURL:    "",
		RefreshRate:    fmt.Sprintf("%d", rs),
		ResetFirmware:  false,
	}
	json.NewEncoder(w).Encode(res)
}
