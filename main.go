package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders(
		"Authorization",
		"Access-Token",
	)
	config.AllowCredentials = true

	r.Use(cors.New(config))

	r.GET("/hellow", func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Access-Token")
		ctx.JSON(200, gin.H{"token": token})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
