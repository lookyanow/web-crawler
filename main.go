package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// type Site struct {
// 	URL string
// }

// type Result struct {
// 	URL    string
// 	Status int
// }

func crawl(url string) int {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error %s, url: %s", err, url)
	}
	time.Sleep(1000 * time.Millisecond)
	return res.StatusCode
}
func main() {
	fmt.Println("Web crawler example program")

	urlList := []string{
		"http://yandex.ru",
		"http://google.com",
		"https://youtube.com",
		"https://ya.ru",
		"https://habr.com",
	}

	for _, url := range urlList {
		status := crawl(url)
		log.Printf("URL [%s] call finished: %d", url, status)
	}
}
