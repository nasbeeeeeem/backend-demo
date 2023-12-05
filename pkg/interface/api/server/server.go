package server

import (
	"backend-demo/pkg/infrastructure/database"
	"backend-demo/pkg/infrastructure/repositoryimpl"
	"backend-demo/pkg/interface/api/handler"
	"backend-demo/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// サーバーの起動処理
func Server(dsn string) {
	// 依存性の注入
	dbClient, err := database.Conn(dsn)
	if err != nil {
		panic(err)
	}
	userRepoImpl := repositoryimpl.NewUserRepoImpl(dbClient)
	userUseCase := usecase.NewUseCase(userRepoImpl)
	userHandler := handler.NewUserHandler(userUseCase)

	eventRepoImpl := repositoryimpl.NewEventRepoImpl(dbClient)
	eventUseCase := usecase.NewEventUseCase(eventRepoImpl)
	eventHandler := handler.NewEventtHandler(eventUseCase)

	r = gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://fetch-api-sample-747dpngw3q-an.a.run.app", "https://fetch-api.haebeal.net"}
	config.AddAllowHeaders(
		"Authorization",
		"Access-Token",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
	)
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// グループ関係のエンドポイント
	groups := r.Group("/groups")
	{
		groups.POST("/")
	}
	// ユーザー関係のエンドポイント
	users := r.Group("/users")
	{
		// 全ユーザの取得
		users.GET("/", userHandler.HandleUsers)
		// 自分の取得
		users.GET("/me", userHandler.HandleMeInfo)
		// ユーザーの更新
		users.PUT("/me", userHandler.HandleUpdate)
		// ユーザーの削除
		users.DELETE("/me", userHandler.HandleDelete)
	}

	// 支払い関係のエンドポイント
	// payment := r.Group("/payment")
	// {
	// 	// 支払い情報（現行でトップページに表示されている情報）
	// 	// 自分に関係ある支払い情報
	// 	payment.GET("/")
	// 	// 新規支払い登録
	// 	payment.POST("/")
	// 	// :idごとの支払い情報
	// 	// 参照
	// 	payment.GET("/:id")
	// 	// 修正
	// 	payment.POST("/:id")
	// 	// 削除
	// 	payment.DELETE("/:id")
	// }

	// イベント関係のエンドポイント
	events := r.Group("/events")
	{
		// 作成
		events.POST("/", eventHandler.HandleCreate)
		// 一覧参照
		events.GET("/", eventHandler.HnadleEvents)
		// ：idごとのイベント
		// 参照
		// event.GET("/:id")
		// 修正
		// event.PUT("/", eventHandler.HandleUpdate)
		// 削除
		// event.DELETE("/", eventHandler.HandleDelete)
	}

	log.Println("Server running...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
