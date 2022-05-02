// handlers.article.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// Получает список топиков
	articles := getAllArticles()

	// Обрабатывает шаблон index.html, передавая ему список топиков
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,

		// Use the index.html template
		"index.html",

		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
