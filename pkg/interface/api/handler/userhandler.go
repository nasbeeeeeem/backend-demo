package handler

import (
	"backend-demo/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	// HandleCreate(c *gin.Context)
	HandleMeInfo(c *gin.Context)
	HandleUsers(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}

type userHandler struct {
	useCase usecase.UserUseCase
}

// handleByEmail implements Handler.
func (h *userHandler) HandleMeInfo(c *gin.Context) {
	type request struct {
		Email string `json:"email"`
	}

	requestBody := new(request)

	user, err := h.useCase.MeInfo(c.Request.Context(), requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":      user.ID,
		"name":        user.Name,
		"email":       user.Email,
		"photoUrl":    user.PhotoURL,
		"accountCode": user.AccountCode,
		"bankCode":    user.BankCode,
		"branchCode":  user.BranchCode,
	})
}

// handleUsers implements Handler.
func (h *userHandler) HandleUsers(c *gin.Context) {
	users, err := h.useCase.Users(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// HandleUpdate implements Handler.
func (h *userHandler) HandleUpdate(c *gin.Context) {
	type request struct {
		ID          int    `json:"id"`
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

	updateUser, err := h.useCase.Update(c, r.ID, r.Name, r.Email, r.PhotoUrl, r.AccountCode, r.BankCode, r.BranchCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"userId":      updateUser.ID,
			"name":        updateUser.Name,
			"email":       updateUser.Email,
			"photoUrl":    updateUser.PhotoURL,
			"accountCode": updateUser.AccountCode,
			"bankCode":    updateUser.BankCode,
			"branchCode":  updateUser.BranchCode,
		})
	}
}

func (h *userHandler) HandleDelete(c *gin.Context) {
	type request struct {
		ID int `json:"id"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	deleteUser, err := h.useCase.Delete(c, r.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"userId":      deleteUser.ID,
			"name":        deleteUser.Name,
			"email":       deleteUser.Email,
			"photoUrl":    deleteUser.PhotoURL,
			"accountCode": deleteUser.AccountCode,
			"bankCode":    deleteUser.BankCode,
			"branchCode":  deleteUser.BranchCode,
		})
	}
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		useCase: userUseCase,
	}
}
