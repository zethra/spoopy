package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download", downloadHandler)
	port := ":8082"
	log.Printf("Starting on port %v", port)
	http.ListenAndServe(port, nil)
}

func downloadHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Bad request", 400)
		log.Fatal(err)
	}
	rawUrl := request.Form.Get("url")
	resp, err := http.Get(rawUrl)
	if err != nil {
		http.Error(writer, "Could not retrive file", 500)
		log.Fatal(err)
		return
	}
	if resp.Header.Get("Content-Disposition") == "" {
		encodedUrl, err := url.Parse(rawUrl)
		if err != nil {
			http.Error(writer, "Could not retrive file", 500)
			log.Fatal(err)
			return
		}
		writer.Header().Set("Content-Disposition", "attachment; filename="+path.Base(encodedUrl.Path)+".html")
	} else {
		writer.Header().Set("Content-Disposition", resp.Header.Get("Content-Disposition")+".html")
	}
	writer.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	writer.Header().Set("Last-Modified", resp.Header.Get("Last-Modified"))
	writer.Header().Set("Content-Type", "text/html")
	io.Copy(writer, resp.Body)
}
