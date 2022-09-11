package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shzxcv/go-restapi/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler.HelloHandler)
	r.HandleFunc("/article", handler.PostArticleHandler)
	r.HandleFunc("/article/list", handler.ArticleListHandler)
	r.HandleFunc("/article/1", handler.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handler.PostNiceHandler)
	r.HandleFunc("/comment", handler.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
