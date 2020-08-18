package main
//
//import (
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//	"log"
//	"math/rand"
//	"net/http"
//	url2 "net/url"
//	"time"
//)
//
//var (
//	//proxies = []string{"socks5://104.236.26.27:38801", "socks5://104.248.63.49:30588", "socks5://104.238.111.150:46470", "socks5://104.238.97.215:15181", "socks5://104.248.63.18:30588", "socks5://167.99.230.114:5577", "socks5://159.89.49.60:31264", "socks5://174.77.111.196:4145", "socks5://174.70.241.17:4145", "socks5://132.148.159.44:13271", "socks5://173.248.156.214:29765", "socks5://174.70.241.18:24404", "socks5://184.168.146.10:1158", "socks5://184.178.172.18:15280", "socks5://184.185.2.12:4145", "socks5://184.185.2.244:4145", "socks5://192.111.129.148:4145", "socks5://184.176.166.13:4145", "socks5://184.178.172.5:15303", "socks5://184.178.172.13:15311", "socks5://47.49.12.169:46451", "socks5://64.118.88.52:9416", "socks5://72.221.164.35:60670", "socks5://47.49.12.165:17326", "socks5://47.49.12.168:24119", "socks5://64.118.87.5:14732", "socks5://47.89.249.147:58629", "socks5://50.62.31.62:20636", "socks5://174.70.241.10:4145", "socks5://167.71.119.18:9050", "socks5://174.77.111.197:4145", "socks5://174.76.48.233:4145", "socks5://174.64.199.82:4145", "socks5://174.70.241.8:24398", "socks5://184.176.166.20:4145"}
//)
//
//func main() {
//	movieLink, err := GetMovieLink("https://reelgood.com/search?q=simpsons")
//	if err != nil {
//		log.Println(err)
//	}
//
//	isOnNetflix, isOnPrime, isOnHulu, isOnDisney, err := GetStreamingDetails("https://reelgood.com" + movieLink)
//	if err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println("Movie Link:")
//	fmt.Println(movieLink)
//
//	fmt.Println("\nOn Netflix?")
//	fmt.Println(isOnNetflix)
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
//	rand.Seed(time.Now().UnixNano())
//
//	//proxyURL, _ := url2.Parse(proxies[rand.Intn(len(proxies))])
//	proxyURL, _ := url2.Parse("socks5://ujigjgxi-US-rotate:8oihb177hbmn@185.30.232.51:1080")
//	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
//
//	res, err := client.Get(url)
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
//	rand.Seed(time.Now().UnixNano())
//
//	proxyURL, _ := url2.Parse("socks5://ujigjgxi-US-rotate:8oihb177hbmn@185.30.232.51:1080")
//	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
//
//	res, err := client.Get(url)
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