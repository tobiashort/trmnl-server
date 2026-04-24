package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LogHandler struct {
}

func (h LogHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.parsePayload(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h LogHandler) parsePayload(payload []byte) error {
	var obj map[string]any
	err := json.Unmarshal(payload, &obj)
	if err != nil {
		return err
	}
	logsArray, ok := obj["logs"].([]any)
	if !ok {
		return fmt.Errorf("logs cast error")
	}
	for _, item := range logsArray {
		itemMap, ok := item.(map[string]any)
		if !ok {
			return fmt.Errorf("item cast error")
		}
		msg, ok := itemMap["message"].(string)
		if !ok {
			return fmt.Errorf("message cast error")
		}
		log.Println(msg)
	}
	return nil
}
