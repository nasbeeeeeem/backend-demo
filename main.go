package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type payload struct {
	Iss            string `json:"iss"`
	Azp            string `json:"azp"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Email          string `json:"email"`
	Email_verified bool   `json:"email_verified"`
	At_hash        string `json:"at_hash"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Locale         string `json:"locale"`
	Iat            int64  `json:"iat"`
	Exp            int64  `json:"exp"`
}

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

	r.GET("/users/me", func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Access-Token")
		jwt := ctx.Request.Header.Get("X-Apigateway-Api-Userinfo")
		dec, err := base64.StdEncoding.DecodeString(jwt)
		if err != nil {
			log.Fatal(err)
		}

		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("auth-token",
			zap.String("token", token),
			zap.String("jwt", jwt),
			zap.String("decode", string(dec)),
		)

		var p payload
		if err := json.Unmarshal(dec, &p); err != nil {
			log.Fatal(err)
		}

		ctx.JSON(200, gin.H{"token": token, "name": p.Name, "email": p.Email})

	})
	// r.GET("/me", handler.HandleMe)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
