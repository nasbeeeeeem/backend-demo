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
	HandleCreate(c *gin.Context)
	HandleUsers(c *gin.Context)
	HandleMeInfo(c *gin.Context)
	HandleUpdate(c *gin.Context)
}

type handler struct {
	useCase usecase.UseCase
}

func (h *handler) HandleCreate(c *gin.Context) {
	type request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	newUser, err := h.useCase.Create(c, r.Name, r.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
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

	token := c.Request.Header.Get("Access-Token")
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
		"token": token,
		"name":  user.Name,
		"email": user.Email,
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

// HandleUpdate implements Handler.
func (h *handler) HandleUpdate(c *gin.Context) {
	type request struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhotoUrl    string `json:"photoUrl"`
		AccountCode string `json:"accountCode"`
		BankCode    string `json:"bankCode"`
		BranchCode  string `json:"branchCode"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := h.useCase.Update(c, r.Name, r.Email, r.PhotoUrl, r.AccountCode, r.BankCode, r.BranchCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully."})
}

func NewHandler(userUseCase usecase.UseCase) Handler {
	return &handler{
		useCase: userUseCase,
	}
}
