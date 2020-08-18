package main

//import (
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/PuerkitoBio/goquery"
//)
//
//func main() {
//	movieLink, err := GetMovieLink("https://reelgood.com/search?q=simpsons")
//	if err != nil {
//		log.Println(err)
//	}
//
//	isOnNetlifx, isOnPrime, isOnHulu, isOnDisney, err := GetStreamingDetails("https://reelgood.com" + movieLink)
//
//	fmt.Println("Movie Link:")
//	fmt.Println(movieLink)
//
//	fmt.Println("\nOn Netflix?")
//	fmt.Println(isOnNetlifx)
//
//	fmt.Println("\nOn Prime?")
//	fmt.Println(isOnPrime)
//
//	fmt.Println("\nOn Hulu?")
//	fmt.Println(isOnHulu)
//
//	fmt.Println("\nOn Disney?")
//	fmt.Println(isOnDisney)
//}
//
//// GetMovieLink gets the href to details page for entered movie/show
//func GetMovieLink(url string) (string, error) {
//	res, err := http.Get(url)
//	if err != nil {
//		return "", err
//	}
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		return "", err
//	}
//
//	var links []string
//	doc.Find(".e1qyeclq5 a").Each(func(i int, s *goquery.Selection) {
//		link, ok := s.Attr("href")
//		if !ok {
//			log.Println("No link found.")
//		}
//
//		links = append(links, link)
//	})
//	return links[0], nil
//}
//
//func GetStreamingDetails(url string) (bool, bool, bool, bool, error) {
//	res, err := http.Get(url)
//	if err != nil {
//		return false, false, false, false, err
//	}
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		return false, false, false, false, err
//	}
//
//	var isOnNetflix bool
//	var isOnHulu bool
//	var isOnPrime bool
//	var isOnDisney bool
//	var text []string
//	doc.Find(".e1udhou113").Each(func(i int, s *goquery.Selection) {
//		text = append(text, s.Text())
//		//isOnNetflix = text[0] == "Netflix"
//
//		for _, v := range text {
//			if v == "Netflix" {
//				isOnNetflix = true
//			}
//
//			if v == "Prime Video" {
//				isOnPrime = true
//			}
//
//			if v == "Hulu" {
//				isOnHulu = true
//			}
//
//			if v == "Disney+" {
//				isOnDisney = true
//			}
//		}
//	})
//
//	return isOnNetflix, isOnPrime, isOnHulu, isOnDisney, err
//}