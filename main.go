package main

import (
	"CRM/database"
	"CRM/handler"
	"CRM/repository"
	"CRM/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()
	r := repository.NewRepositoryAdapter(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHanerhandlerAdapter(s)

	router := gin.Default()
	router.GET("/api/get", h.GetHan)
	router.POST("/add", h.AddHan)

	err := router.Run(":9001")
	if err != nil {
		panic(err.Error())
	}
}
