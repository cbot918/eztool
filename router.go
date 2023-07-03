package main

import (
	"github.com/cbot918/eztool/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, dep *Dep) *gin.Engine {
	ctr := controller.NewController(dep.Db, dep.Cache)

	r.GET("/ping", ctr.Ping)
	r.GET("/health", ctr.Health)
	r.POST("/auth/signin")
	r.POST("/auth/signup", ctr.Signup)
	r.POST("/install", ctr.Install)
	return r
}
