package pkg

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func isEztoolExisted() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("os.UserHomeDir failed")
		panic(err)
	}
	return slices.Contains(Ls(home), ".eztool")
}

func PreCheckEnv() {
	fmt.Println(isEztoolExisted())
}

func PreSetupEnv() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("os.UserHomeDir failed")
		panic(err)
	}
	Mkdir(fmt.Sprintf("%s/.eztool", home))
	Mkdir(fmt.Sprintf("%s/.eztool/bin", home))
	Mkdir(fmt.Sprintf("%s/.eztool/temp", home))
}
