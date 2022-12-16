package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	// `json:`(注意json後面不能有空格)
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	// 指定回傳的標頭為 json
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go Blog"
	indexData.Desc = "first"
	jsonStr, _ := json.Marshal(indexData)
	// w.Write([]byte("hello go blog"))
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go Blog"
	indexData.Desc = "first"
	t := template.New("index.html")
	// 解析檔案
	path, _ := os.Getwd()
	// 因為首頁有多個template嵌套, 所以需要將涉及到的template都進行解析
	home := path + "/template/home.html"
	header := path + "/template/header.html"
	footer := path + "/template/footer.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer)
	t.Execute(w, indexData)

}

func main() {

	// 設定路徑 server
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	// 響應 request
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
