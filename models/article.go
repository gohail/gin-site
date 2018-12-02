package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewArticle(ID int, title string, content string) Article {
	return Article{ID: ID, Title: title, Content: content}
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []Article {
	return articleList
}

func GetArticleByID(ID int) (Article, error) {
	for _, val := range articleList {
		if val.ID == ID {
			return val, nil
		}
	}
	return Article{}, errors.New("Article not found")
}
