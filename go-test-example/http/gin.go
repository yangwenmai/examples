package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) { c.String(http.StatusOK, "Hello") })
	engine.GET("/world", func(c *gin.Context) { c.String(http.StatusOK, "world") })

	req := httptest.NewRequest(http.MethodGet, "/world", nil)
	w := httptest.NewRecorder()

	engine.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
}
