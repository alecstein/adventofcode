package utils

import (
	"io"
	"log"
	"net/http"
)

func GetPuzzleInput(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:120.0) Gecko/20100101 Firefox/120.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "session=53616c7465645f5f4d886fc038ad5d8ec590b91a029eec7512efcdbd9344bbb378a078aab12a11c8b34e947f45c9a2daa16b7f99db66c0da56207062d3d83b0e")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText), err
}

func ReverseString(s string) string {
	reversed := make([]rune, len(s))
	for i, r := range s {
		reversed[len(s)-1-i] = r
	}
	return string(reversed)
}
