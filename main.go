package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(url string) int {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error %s, url: %s", err, url)
	}
	time.Sleep(1000 * time.Millisecond)
	return res.StatusCode
}

func worker(id int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("[%d] starting crawl url %s", id, site.URL)

		res, err := http.Get(site.URL)
		if err != nil {
			log.Printf("[%d] Error %s, url: [%s]", id, err, site.URL)
		}
		time.Sleep(1000 * time.Millisecond)
		results <- Result{
			URL:    site.URL,
			Status: res.StatusCode,
		}
	}
}

func main() {
	fmt.Println("Web crawler example program(with worker pool)")

	// Making example URL for our crawler app
	urlList := []string{
		"http://yandex.ru",
		"http://google.com",
		"https://youtube.com",
		"https://ya.ru",
		"https://habr.com",
		"https://dsfdsf.ru",
		"https://kubernetes.io",
		"https://lenta.ru",
	}

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for _, url := range urlList {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		result := <-results
		log.Println(result)
	}

}
