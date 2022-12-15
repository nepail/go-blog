package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func main() {

	// 設定路徑 server
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	// 響應 request
	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
