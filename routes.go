package main

import "github.com/gin-gonic/gin"

func initializeRoutes() {
	router := gin.Default()

	// определение роута главной страницы
	router.GET("/", showIndexPage)
}
