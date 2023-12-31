package utils

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// From curlconverter.com
// Return the puzzle input as a slice of strings
func GetPuzzleInput(url string) ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", SessionCookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bodyText), "\n")
	// drop any empty lines
	// for i, line := range lines {
	// 	if line == "" {
	// 		lines = append(lines[:i], lines[i+1:]...)
	// 	}
	// }
	return lines, err
}

// Open file and get puzzle input as lines of strings
func GetPuzzleInputFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
