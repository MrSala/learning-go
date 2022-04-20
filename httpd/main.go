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

	r := chi.NewRouter()

	r.Get("/articles", handler.ArticlesGet(feed))
	r.Post("/articles", handler.ArticlesPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
