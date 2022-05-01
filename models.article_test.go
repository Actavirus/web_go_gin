// models.article_test.go
// В этом тесте используется функция getAllArticles() для получения
// списка всех топиков. Сперва этот тест проверяет, что эта функция
// получает список топиков и этот список идентичен списку,
// загруженному в глобальную переменную articleList. Затем он
// проходит в цикле по списку топиков для проверки уникальности
// каждого. Если хотя бы одна из этих проверок не удалась, тест
// возвращает неудачу.

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the lenght of the list of articles returned is the
	// same as the lenght of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
