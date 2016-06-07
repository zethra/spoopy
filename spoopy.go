package main

import (
	"net/http"
	"log"
	"io"
	//"net/url"
	//"path"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8082", nil)
}

func downloadHandler(writer http.ResponseWriter, request *http.Request)  {
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
	/*if resp.Header.Get("Content-Disposition") == "" {
		encodedUrl, err := url.Parse(rawUrl)
		if err != nil {
			http.Error(writer, "Could not retrive file", 500)
			log.Fatal(err)
			return
		}
		writer.Header().Set("Content-Disposition", "attachment; filename=" + path.Base(encodedUrl.Path) + ".html")
	} else {
		writer.Header().Set("Content-Disposition", resp.Header.Get("Content-Disposition") + ".html")
	}*/
	writer.Header().Set("Content-Type", "text/html")
	io.Copy(writer, resp.Body)
}