package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
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

func HandleMe(c *gin.Context) {
	jwt := c.Request.Header.Get("X-Apigateway-Api-Userinfo")
	dec, err := base64.RawURLEncoding.DecodeString(jwt)
	if err != nil {
		log.Fatal(err)
	}

	var p payload

	if err := json.Unmarshal(dec, &p); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"name": p.Name, "email": p.Email})
}
