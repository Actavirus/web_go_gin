package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Со­зда­ём ро­у­тер(роутер по-умолчанию в Gin)
	router := gin.Default()

	// За­гру­жа­ем шаб­ло­ны из пап­ки templates.
	// Обработаем шаблоны вначале, так что их не нужно будет перечитывать
	// ещё раз. Из-за этого вывод HTML-страниц такой быстрый.
	router.LoadHTMLGlob("templates/*")

	// // Инициализируем роуты
	// initializeRoutes()

	// определение роута главной страницы
	router.GET("/", showIndexPage)

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	router.GET("/article/view/:article_id", getArticle)

	// За­пуск при­ло­же­ния (При­ло­же­ние за­пу­стит­ся на localhost и 8080 пор­те, по-умол­ча­нию.)
	router.Run()
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
