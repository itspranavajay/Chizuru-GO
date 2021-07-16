package bot

import (
	"log"
	"net/http"
)

func getType(url string) string {
	resp, err := http.Get(url)
	if resp == nil {
		return "FAILED"
	}
	if err != nil {
		log.Println(err.Error())
	}

	content_type := resp.Header.Get("Content-Type")
	return content_type
}
