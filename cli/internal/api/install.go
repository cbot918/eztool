package api

import (
	"cli/internal/api/service"
	u "cli/internal/util"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type InstallRequest struct {
	Name string `json:"name" `
}

type InstallResponse struct {
	Message string `json:"message"`
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
	service := service.NewService()
	service.Install()
}
