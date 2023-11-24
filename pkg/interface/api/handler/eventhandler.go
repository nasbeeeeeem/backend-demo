package handler

import (
	"backend-demo/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler interface {
	HandleCreate(c *gin.Context)
	HnadleEvents(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}

type eventHandler struct {
	eventUsecase usecase.EventUseCase
}

// HandleCreate implements EventHandler.
func (h *eventHandler) HandleCreate(c *gin.Context) {
	type request struct {
		Name   string `json:"name"`
		UserId int    `json:"userId"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	newEvent, err := h.eventUsecase.Create(c, r.Name, r.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"event": newEvent})
	}
}

// HandleDelete implements EventHandler.
func (h *eventHandler) HandleDelete(c *gin.Context) {
	type request struct {
		EventId int `json:"eventId"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	deleteEvent, err := h.eventUsecase.Delete(c, r.EventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"deleteEvent": deleteEvent})
	}
}

// HandleUpdate implements EventHandler.
func (h *eventHandler) HandleUpdate(c *gin.Context) {
	type request struct {
		EventId int `json:"eventId"`
	}

	r := new(request)
	if err := c.Bind(&r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	updateEvent, err := h.eventUsecase.Update(c, r.EventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"updateEvent": updateEvent})
	}
}

// HnadleEvents implements EventHandler.
func (h *eventHandler) HnadleEvents(c *gin.Context) {
	events, err := h.eventUsecase.Events(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"events": events})
	}
}

func NewEventtHandler(eventUsecase usecase.EventUseCase) EventHandler {
	return &eventHandler{
		eventUsecase: eventUsecase,
	}
}
