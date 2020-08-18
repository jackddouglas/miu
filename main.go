package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	url2 "net/url"
	"time"
)

func main() {
	movieLink, err := GetMovieLink("https://reelgood.com/search?q=limitless")
	if err != nil {
		log.Println(err)
	}

	title, description, isOnNetflix, isOnPrime, isOnHulu, isOnDisney, err := GetStreamingDetails("https://reelgood.com" + movieLink)
	if err != nil {
		log.Println(err)
	}

	moviePrice, err := GetAmazonPrice("https://www.amazon.com/s?k=" + title + "&i=movies-tv")

	fmt.Println(moviePrice)

	fmt.Println(title)
	fmt.Println("\n" + description)

	fmt.Println("\nWhere to get it:")

	if isOnNetflix {
		fmt.Println(Pad("Netflix", "subscription"))
	}

	if isOnPrime {
		fmt.Println(Pad("Prime Video", "subscription"))
	}

	if isOnHulu {
		fmt.Println(Pad("Hulu", "subscription"))
	}

	if isOnDisney {
		fmt.Println(Pad("Disney+", "subscription"))
	}
}

// GetMovieLink gets the href to details page for entered movie/show
func GetMovieLink(url string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	proxyURL, _ := url2.Parse("socks5://ujigjgxi-US-rotate:8oihb177hbmn@185.30.232.51:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	res, err := client.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var links []string
	doc.Find(".e1qyeclq5 a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			log.Println("No link found.")
		}

		links = append(links, link)
	})
	return links[0], nil
}

// GetStreamingDetails scrapes streaming data from details page for movie/show
func GetStreamingDetails(url string) (string, string, bool, bool, bool, bool, error) {
	rand.Seed(time.Now().UnixNano())

	proxyURL, _ := url2.Parse("socks5://ujigjgxi-US-rotate:8oihb177hbmn@185.30.232.51:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	res, err := client.Get(url)
	if err != nil {
		return "Error", "Error", false, false, false, false, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "Error", "Error", false, false, false, false, err
	}

	var title string
	doc.Find(".e14injhv7").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})

	var description string
	doc.Find(".e50tfam1 p").Each(func(i int, s *goquery.Selection) {
		description = s.Text()
	})

	var isOnNetflix bool
	var isOnHulu bool
	var isOnPrime bool
	var isOnDisney bool
	var text []string
	doc.Find(".e1udhou113").Each(func(i int, s *goquery.Selection) {
		text = append(text, s.Text())

		for _, v := range text {
			if v == "Netflix" {
				isOnNetflix = true
			}

			if v == "Prime Video" {
				isOnPrime = true
			}

			if v == "Hulu" {
				isOnHulu = true
			}

			if v == "Disney+" {
				isOnDisney = true
			}
		}
	})

	return title, description, isOnNetflix, isOnPrime, isOnHulu, isOnDisney, err
}

// GetAmazonPrice finds the price of the movie/show and returns it
func GetAmazonPrice(url string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	proxyURL, _ := url2.Parse("socks5://ujigjgxi-US-rotate:8oihb177hbmn@185.30.232.51:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "Error", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:10.0) Gecko/20100101 Firefox/10.0")

	res, err := client.Do(req)
	if err != nil {
		return "Error", err
	}

	if res.StatusCode > 500 {
		fmt.Println("So uh, there was a problem. We got blocked.")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "Error", err
	}

	//fmt.Println(doc.Contents().Text())

	var prices []string
	doc.Find(".a-offscreen").Each(func(i int, s *goquery.Selection) {
		prices = append(prices, s.Text())
		fmt.Println(s.Text())
	})

	return prices[1], err
}

// Pad adds padding functionality to formatted lines
func Pad(left string, right string) string {
	for len(left + right) < 125 {
		left += " • "
	}

	return left + right
}