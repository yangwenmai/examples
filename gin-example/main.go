package main

import (
	"github.com/prometheus/common/log"
	"math/rand"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func init() {
	log.Warnln("init...")
	// rand.Seed(time.Now().UnixNano())
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	
	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		rs:=[]string{"3", "1", "4", ""}
		x:=rand.Intn(len(rs))
		ret := rs[x]
		c.JSON(http.StatusOK, gin.H{"user": user, "value": ret})
	})
	
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}