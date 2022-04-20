package main

import (
	"fmt"
	"learning-go/handler"
	"learning-go/platform/articles"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"
	feed := articles.New()
	feed.Add(articles.Article{
		Title:   "Learning Go",
		Summary: "Learning Go to be able to work in ispec and won't become a neck",
		Author:  "Martin",
	})

	r := chi.NewRouter()

	r.Get("/articles", handler.ArticlesGet(feed))
	r.Post("/articles", handler.ArticlesPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
