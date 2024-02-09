package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Mondoo Engineer!")
	})
	return r
}

func main() {
	r := setupRouter()
	env_port := os.Getenv("PORT")
	if env_port == "" {
		env_port = ":29090"
	} else {
		env_port = ":" + env_port
	}

	r.Run(env_port)
}
