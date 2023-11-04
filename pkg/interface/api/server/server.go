package server

import (
	"backend-demo/pkg/interface/api/handler"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Server() {
	r = gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://fetch-api-sample-747dpngw3q-an.a.run.app/"}
	config.AddAllowHeaders(
		"Authorization",
		"Access-Token",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
	)
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// ユーザー関係のエンドポイント
	users := r.Group("/users")
	{
		// 全ユーザの取得
		users.GET("/")
		// 自分の取得
		users.GET("/me", handler.HandleMe)
	}

	// 支払い関係のエンドポイント
	payment := r.Group("/payment")
	{
		// 支払い情報（現行でトップページに表示されている情報）
		// 自分に関係ある支払い情報
		payment.GET("/")
		// 新規支払い登録
		payment.POST("/")
		// :idごとの支払い情報
		// 参照
		payment.GET("/:id")
		// 修正
		payment.POST("/:id")
		// 削除
		payment.DELETE("/:id")
	}

	// イベント関係のエンドポインおｔ
	event := r.Group("/event")
	{
		// イベント一覧
		event.GET("/")
		// ：idごとのイベント
		// 参照
		event.GET("/:id")
		// 修正
		event.PUT("/:id")
		// 削除
		event.DELETE("/:id")
	}

	log.Println("Server running...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
