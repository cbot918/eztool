package api

import (
	"fmt"
	"net/http"
)

var log = fmt.Println

func Login() {
	res, err := http.Get("http://localhost:5500/ping")
	if err != nil {
		log("http.Get failed")
		return
	}
	log(res)
}
