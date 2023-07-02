package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	log  = fmt.Println
	logf = fmt.Printf
	logj = func(obj interface{}) {
		bytes, _ := json.MarshalIndent(obj, "", "  ")
		fmt.Println(string(bytes))
	}
)

type Controller struct {
	Db    *sql.DB
	Cache *redis.Client
}

func NewController(db *sql.DB, cache *redis.Client) *Controller {
	return &Controller{
		Db:    db,
		Cache: cache,
	}
}

func (ctr *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ctr *Controller) Health(c *gin.Context) {
	msg := ""
	if err := ctr.Db.Ping(); err != nil {
		msg += "db ping failed"
	}
	if res := ctr.Cache.Ping(context.Background()); res.Val() != "PONG" {
		msg += "cache ping failed"
	}
	if msg == "" {
		msg += "db ok cache ok"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

type SignupParam struct {
	A string `json:"a"`
}

type SignupResponse struct {
}

func (ctr *Controller) Signup(c *gin.Context) {
	req := &SignupParam{}
	if err := c.BindJSON(&req); err != nil {
		log("c.BindJSON failed")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong request body",
		})
		return
	}
	logj(req)
	c.JSON(http.StatusOK, gin.H{
		"message": "signup",
	})
	//
}
