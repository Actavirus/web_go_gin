// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "article 1 body"},
	{ID: 2, Title: "Article 2", Content: "article 2 body"},
	{ID: 3, Title: "Article 3", Content: "article 3 body"},
	{ID: 4, Title: "Article 4", Content: "article 4 body"},
}

// Return a list of all articles.
func getAllArticles() []article {
	return articleList
}

// Эта функция считывает список топиков в цикле и возвращает
// топик, ID которого соответствует переданному ID. Если такого
// топика нет, она возвращает ошибку.
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
