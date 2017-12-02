// task5
package main

import (
	"encoding/json"
	"net/http"
)

var decompressedUrls map[string]string

func getCompressedUrl(url string) string {
	compressedUrl := ""
	i := len(decompressedUrls) + 1
	for ; i > 0; i /= 26 {
		compressedUrl += string('a' + i%26)
	}
	decompressedUrls[compressedUrl] = url
	return compressedUrl
}

type UserRequest struct {
	Url string `json:"url"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var ureq UserRequest

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&ureq)

		if err != nil || ureq.Url == "" {
			http.Error(w, "", 400)
			return
		}

		responseMap := make(map[string]string)
		responseMap["key"] = getCompressedUrl(ureq.Url)
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	} else if r.Method == "GET" {
		compressedUrl := r.RequestURI[1:]
		decompressedUrl, exists := decompressedUrls[compressedUrl]
		if !exists {
			http.Error(w, "not found", 404)
			return
		}
		http.Redirect(w, r, decompressedUrl, 301)
	}
}

func main() {
	decompressedUrls = make(map[string]string)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
