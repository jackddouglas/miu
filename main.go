package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	url2 "net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	movieLink, err := GetMovieLink("https://reelgood.com/search?q=bill%20gates")
	if err != nil {
		log.Println(err)
	}

	title, description, isOnNetflix, isOnPrime, isOnHulu, isOnDisney, err := GetStreamingDetails("https://reelgood.com" + movieLink)
	if err != nil {
		log.Println(err)
	}

	rentPrice, err := GetAmazonPrice("https://www.amazon.com/s?k=" + title + "&i=movies-tv&ref=nb_sb_noss_1")

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

	fmt.Println(Pad("Amazon (rent)", rentPrice))
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

	proxyURL, _ := url2.Parse("http://114.239.171.181:4216")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "Error", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Cookie", "aws-ubid-main=656-2461571-6267250; aws-session-id=132-3199980-3265813; aws-session-id-time=2227936015l; csm-hit=tb:EZA90K41RH7TPYY2T8Y3+b-NSHGSSSE64MJER40A9CX|1597892738291&t:1597892738291&adb:adblk_no; aws-session-token=\"S1PRmUPtAXqce+Id2xIstoU7qf07Xsaf9V/2Nzy7snBw30aVYKO7iU824l4cdHSiray8NiC1vWNtGLnkeEuda8NzM4lENMTn8/QjhlBZ+M72XoWeMQ+zM9LuerXD+qpYLPQCvqV/yifXj+6JBTsbntx5x/2LOdBvuOhl1iMwXqghvJqxalZGvUNqLQGwzVAMMmCbTKJqiwYpTIZ9a8aUkdQM6NzOXA5ySoA6YltnC2s=\"; aws-x-main=\"895o@GGNHaCyn0ZrcAxZ17?rm2DktswOzt98aj6…t-main=\"EoMCkUfeFtJyk6Jc/Z8+ICmd3lmgZN9TcJFfWONpvNk=\"; sst-main=Sst1|PQGHZ9cjWrhIwxC6pG1jHbdMC1oqShS_0KZ2NPTVGVfLXHS8XsidXC_rh7Z1qeC9VAc09z_GXgoj5YbUES6MYUTN1cwlabhmo28JV7OGGBKIMsoXUFBXpue404SzcU9YMF5i8TXNqcgcVOvUaH5JksenLQIyXy4xh7UUkN7ThQwZ7dIc0MFzpi1FdcTD1CYI2X2XynxFdZjCenZCvDffFnrm2Tmgq4FG_t_ncllnswCRAOeT-sj1_ExVCYdjc-TXm3Q8lmsGd4nBWdy_YHiMHdjN8mNarq99-Y6O7FOqdWU-trL0XyKIIPOqIlcI2hAnW7Cbt1eVn8bVjSb7YzxhJgwHXw; lc-main=en_US; i18n-prefs=USD; lc-main-av=en_US; ubid-main-av=135-2193274-6489920; skin=noskin")

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

	fmt.Println(doc.Contents().Text())

	var prices []string
	doc.Find(".a-offscreen").Each(func(i int, s *goquery.Selection) {
		prices = append(prices, s.Text())
	})

	return prices[1], err
}

// Pad adds padding functionality to formatted lines
func Pad(left string, right string) string {
	for len(left+right) < 125 {
		left += " • "
	}

	return left + right
}
