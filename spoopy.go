package main

import (
	"net/http"
	"log"
	"io"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8080", nil)
}

func downloadHandler(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Bad request", 400)
		log.Fatal(err)
	}
	url := request.Form.Get("url")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(writer, "Could not retrive file", 500)
		log.Fatal(err)
	}
	log.Println(resp.Header.Get("Content-Disposition"))
	writer.Header().Set("Content-Disposition", resp.Header.Get("Content-Disposition") + ".html")
	writer.Header().Set("Content-Type", "text/html")
	io.Copy(writer, resp.Body)
}