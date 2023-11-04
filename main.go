package main

import (
	"backend-demo/pkg/handler"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("server running...")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AddAllowHeaders(
		"Authorization",
		"Access-Token",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
	)
	config.AllowCredentials = true

	r.Use(cors.New(config))

	r.GET("/hello", func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Access-Token")
		jwt := ctx.Request.Header.Get("X-Apigateway-Api-Userinfo")
		log.Println(token)
		log.Println(jwt)
		ctx.JSON(200, gin.H{"token": token})
	})
	r.GET("/users/me", handler.HandleMe)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
