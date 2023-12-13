package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	downloadSite(url)
}

func downloadSite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: HTTP status code %d\n", resp.StatusCode)
		return
	}

	baseURL := getBaseURL(url)
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "a" || token.Data == "img" || token.Data == "link" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						link := attr.Val
						if !strings.HasPrefix(link, "http") {
							link = baseURL + link
						}
						saveFile(baseURL, link)
					}
				}
			}
		}
	}
}

func saveFile(baseUrl string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	fileName := filepath.Base(url)
	file, err := os.Create(baseUrl + "/" + fileName + ".html")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Downloaded:", url)
}

func getBaseURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		return url
	}
	return parts[0] + "//" + parts[2]
}
