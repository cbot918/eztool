package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/exp/slices"
)

func Contains(who []string, target string) bool {
	return slices.Contains(who, target)
}

func GetHome() (home string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("os.UserHomeDir failed")
		panic(err)
	}
	return
}

func GetExt(file string) (ext string) {
	if filepath.Ext(file) == "" && file == "Makefile" {
		ext = "Makefile"
	} else if filepath.Ext(file) == ".sh" {
		ext = ".sh"
	} else {
		ext = filepath.Ext(file)
	}
	return
}

func Ls(path string) (res []string) {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		res = append(res, e.Name())
	}
	return res
}

func Mkdir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Println("os.Mkdir failed")
		panic(err)
	}
}

func Cd(name string) {
	err := os.Chdir(name)
	if err != nil {
		fmt.Println("os.Chdir failed")
		panic(err)
	}
}

func Exec(cmd string) (err error) {
	if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
		return
	}
	return nil
}

func ExecO(cmd string) (string, error) {
	byts, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(byts), nil
}

// func main() {
// 	fmt.Println(Ls("./"))
// }
