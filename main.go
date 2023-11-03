package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Token",
			"Access-Control-Allow-Origin",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/hellow", func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Access-Token")
		ctx.JSON(200, gin.H{"token": token})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
