package api

type Api struct{}

type InstallConfig struct {
	Name   string
	Auth   bool
	Target string
}

func NewApi() *Api {
	return &Api{}
}
