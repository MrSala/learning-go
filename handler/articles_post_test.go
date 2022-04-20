package handler

import (
	"learning-go/platform/articles"
	"learning-go/platform/mock_http"
	"net/http"
	"testing"
)

func TestArticlesPost(t *testing.T) {
	feed := articles.New()

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"title":   "New title",
		"summary": "New summary",
		"author":  "New author",
	})

	handler := ArticlesPost(feed)
	handler(w, r)

	result := w.GetBodyString()

	if result != "Article added" {
		t.Errorf("Handler did not complete")
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Article did not add")
	}

	if feed.GetAll()[0].Title != "New title" {
		t.Errorf("Article bad")
	}
}
