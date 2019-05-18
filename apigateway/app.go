package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	registerOrderService(r)
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

type Message struct {
	Message string `json:"message"`
}

func registerOrderService(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		body := &Message{"Hello, world!"}
		c.JSON(http.StatusOK, body)
	})
}
