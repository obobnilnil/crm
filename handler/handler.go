package handler

import (
	"CRM/model"
	"CRM/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	GetHan(c *gin.Context)
	AddHan(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHanerhandlerAdapter(s service.ServicePort) HandlerPort {
	return handlerAdapter{s: s}
}
func (h handlerAdapter) GetHan(c *gin.Context) {
	data, err := h.s.GetSer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h handlerAdapter) AddHan(c *gin.Context) {
	var req model.Addrequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastID, err := h.s.Addser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully", "lastID": lastID})
}
