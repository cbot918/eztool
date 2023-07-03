package api

import (
	"cli/internal/api/service"
	u "cli/internal/util"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type InstallRequest struct {
	Name   string `json:"name" `
	Token  string `json:"token"`
	Target string `json:"target"`
}

func Install() {
	url := "http://localhost:5500/install"
	client := resty.New()

	data, err := json.Marshal(&InstallRequest{
		Name: "yale",
	})
	if err != nil {
		log("json.Marshal failed")
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	if err != nil {
		log("client request failed")
		return
	}
	if resp.StatusCode() != 200 {
		u.Log("cli stop")
		u.Log(resp)
		return
	}
	// mock response from eztool service
	cfg := &service.InstallResponse{
		Message: "mock response",
		Name:    "yale918",
		Token:   "12345",
		Target:  "gob",
	}
	svc := service.NewService()
	svc.Install(cfg)
}
