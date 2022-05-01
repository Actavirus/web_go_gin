package main

func main() {

}

func initializeRoutes() {
	// определение роута главной страницы
	router.GET("/", showIndexPage)
}
