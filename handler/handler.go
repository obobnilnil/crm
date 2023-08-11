package handler

import (
	"followPtong/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	GetHan(c *gin.Context)
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
