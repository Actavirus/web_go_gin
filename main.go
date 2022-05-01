package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Со­зда­ём ро­у­тер(роутер по-умолчанию в Gin)
	router := gin.Default()

	// За­гру­жа­ем шаб­ло­ны из пап­ки templates.
	// Обработаем шаблоны вначале, так что их не нужно будет перечитывать
	// ещё раз. Из-за этого вывод HTML-страниц такой быстрый.
	router.LoadHTMLGlob("templates/*")


	// Инициализируем роуты
	initializeRoutes()


	// // За­да­ём об­ра­бот­чик ро­у­тов
	// router.GET("/", func(c *gin.Context) {
	// 	// Call the HTML method of the Context to render a template
	// 	c.HTML(
	// 		// Set the HTTP status to 200 (OK)
	// 		http.StatusOK,
	// 		// Use the index.html template
	// 		"index.html",
	// 		// Pass the data that the page uses (in this case, 'title')
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	// За­пуск при­ло­же­ния (При­ло­же­ние за­пу­стит­ся на localhost и 8080 пор­те, по-умол­ча­нию.)
	router.Run()
}
