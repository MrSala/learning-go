package handler

import (
	"encoding/json"
	"learning-go/platform/articles"
	"net/http"
)

func ArticlesPost(feed articles.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(articles.Article{
			Title:   request["title"],
			Summary: request["summary"],
			Author:  request["author"],
		})

		w.Write([]byte("Article added"))
	}
}
