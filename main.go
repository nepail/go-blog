package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go Blog"
	indexData.Desc = "first"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
