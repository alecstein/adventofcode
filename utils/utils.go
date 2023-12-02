package utils

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// From curlconverter.com
func GetPuzzleInput(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", "session=53616c7465645f5f4d886fc038ad5d8ec590b91a029eec7512efcdbd9344bbb378a078aab12a11c8b34e947f45c9a2daa16b7f99db66c0da56207062d3d83b0e")
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

func GetPuzzleInputAsLines(url string) ([]string, error) {
	input, err := GetPuzzleInput(url)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(input, "\n")
	// drop any empty lines
	for i, line := range lines {
		if line == "" {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	return lines, err
}

func Reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}
