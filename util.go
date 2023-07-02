package main

import (
	"encoding/json"
	"fmt"
)

var (
	log  = fmt.Println
	logf = fmt.Printf
	logj = func(obj interface{}) {
		bytes, _ := json.MarshalIndent(obj, "", "  ")
		fmt.Println(string(bytes))
	}
)
