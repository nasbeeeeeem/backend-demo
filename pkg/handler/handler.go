package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleMe(c *gin.Context) {
	// base64でエンコードされたjwtを取得
	jwt := c.Request.Header.Get("X-Apigateway-Api-Userinfo")
	log.Println(jwt)
	c.JSON(200, gin.H{"jwt": jwt})
}
