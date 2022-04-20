package handler

import (
	"learning-go/platform/articles"
	"learning-go/platform/mock_http"
	"net/http"
	"testing"
)

func TestArticlesGet(t *testing.T) {
	feed := articles.New()
	feed.Add(articles.Article{Title: "Title", Summary: "Summary", Author: "Author"})

	handler := ArticlesGet(feed)

	w := &mock_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()

	if len(result) != 1 {
		t.Errorf("Article was not added")
	}

	if result[0]["title"] != "Title" {
		t.Errorf("Article title was not properly set")
	}
}
