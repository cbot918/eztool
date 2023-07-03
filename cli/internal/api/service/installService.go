package service

import (
	"cli/pkg"
	"fmt"
	"log"

	"github.com/cbot918/liby/util"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Install() {
	name := "yale918"
	auth := true
	installTarget := "gob"

	if !auth {
		fmt.Println("驗證未通過")
		return
	}
	path := pkg.GetHome() + "/.eztool/temp/tmp" + util.GetUuid()
	pkg.Mkdir(path)
	pkg.Cd(path)

	murl := "http://localhost:8100" + "/" + name + "/" + installTarget + "/" + "Makefile"
	err := pkg.Downloader(murl)
	if err != nil {
		fmt.Println("pkg.Downloader failed in Install")
		log.Fatal(err)
	}
	if err = pkg.Exec("make"); err != nil {
		fmt.Println("pkg.Exec make failed in Install")
		log.Fatal(err)
	}
	if !pkg.Contains(pkg.Ls(path), installTarget) {
		fmt.Println("something wrong")
	}
	fmt.Println("\ngob was downloaded ")

	// var execstr string
	// switch pkg.GetExt(file) {
	// case ".sh":
	// 	{
	// 		execstr = fmt.Sprintf("./%s", file)
	// 	}
	// case "Makefile":
	// 	{
	// 		execstr = "make"
	// 	}
	// default:
	// 	fmt.Println("other file ext")
	// }

	// // res, err := exec.Command("bash", "-c", ".eztool/scripts/install.sh").Output()
	// res, err := exec.Command("sh", "-c", execstr).Output()
	// if err != nil {
	// 	log("exec.Command failed")
	// 	panic(err)
	// }
	// util.Log(string(res))

}
