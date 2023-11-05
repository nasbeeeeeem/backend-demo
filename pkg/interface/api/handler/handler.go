package handler

import (
	"backend-demo/pkg/usecase"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandleUsers(c *gin.Context)
	HandleMeInfo(c *gin.Context)
}

type handler struct {
	useCase usecase.UseCase
}

// handleByEmail implements Handler.
func (h *handler) HandleMeInfo(c *gin.Context) {
	type request struct {
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

	requestBody := new(request)

	jwt := c.Request.Header.Get("X-Apigateway-Api-Userinfo")
	dec, err := base64.RawURLEncoding.DecodeString(jwt)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(dec, &requestBody); err != nil {
		log.Fatal(err)
	}

	user, err := h.useCase.MeInfo(c.Request.Context(), requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"Email": user.Email,
		"Name":  user.Name,
	})
}

// handleUsers implements Handler.
func (h *handler) HandleUsers(c *gin.Context) {
	users, err := h.useCase.Users(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func NewHandler(userUseCase usecase.UseCase) Handler {
	return &handler{
		useCase: userUseCase,
	}
}