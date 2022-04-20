package handler

import (
	"encoding/json"
	"learning-go/platform/articles"
	"net/http"
)

func ArticlesGet(feed articles.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
