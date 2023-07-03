package util

import (
	"encoding/json"
	"fmt"
)

func Log(a ...any) {
	fmt.Println(a)
}

func Logf(format string, a ...any) {
	fmt.Printf(format, a)
}

func Logj(a ...any) {
	res, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		Log("json.MarshalIndent failed")
		return
	}
	Log(string(res))
}
