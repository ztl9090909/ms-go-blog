package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "码神之路go博客"
	indexData.Desc = "现在是入门教程"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

//func indexHtml(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var indexData IndexData
//	indexData.Title = "码神之路go博客"
//	indexData.Desc = "现在是入门教程"
//	t := template.New("index.html")
//	//1. 拿到当前的路径
//	path, _ := os.Getwd()
//	t1, _ := t.ParseFiles(path + "/template/index.html")
//	t1.Execute(w, indexData)
//}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	viewPath, _ := os.Getwd()
	t, _ = t.ParseFiles(viewPath + "/template/index.html")
	var indexData IndexData
	indexData.Title = "码神之路go博客"
	indexData.Desc = "现在是入门教程"
	fmt.Println(t)
	err := t.Execute(w, indexData)
	fmt.Println(err)
}

func main() {
	// 程序入口，一个项目 只能有一个入口
	// web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
