package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shzxcv/go-restapi/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handler.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handler.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handler.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handler.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handler.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
