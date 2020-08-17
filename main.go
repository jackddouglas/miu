package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0"),
		colly.AllowURLRevisit(),
	)
	c.IgnoreRobotsTxt = false

	rp, err := proxy.RoundRobinProxySwitcher("socks5://104.236.26.27:38801", "socks5://104.248.63.49:30588", "socks5://104.238.111.150:46470", "socks5://104.238.97.215:15181", "socks5://104.248.63.18:30588", "socks5://167.99.230.114:5577", "socks5://159.89.49.60:31264", "socks5://174.77.111.196:4145", "socks5://174.70.241.17:4145", "socks5://132.148.159.44:13271", "socks5://173.248.156.214:29765", "socks5://174.70.241.18:24404", "socks5://184.168.146.10:1158", "socks5://184.178.172.18:15280", "socks5://184.185.2.12:4145", "socks5://184.185.2.244:4145", "socks5://192.111.129.148:4145", "socks5://184.176.166.13:4145", "socks5://184.178.172.5:15303", "socks5://184.178.172.13:15311", "socks5://47.49.12.169:46451", "socks5://64.118.88.52:9416", "socks5://72.221.164.35:60670", "socks5://47.49.12.165:17326", "socks5://47.49.12.168:24119", "socks5://64.118.87.5:14732", "socks5://47.89.249.147:58629", "socks5://50.62.31.62:20636", "socks5://174.70.241.10:4145", "socks5://167.71.119.18:9050", "socks5://174.77.111.197:4145", "socks5://174.76.48.233:4145", "socks5://174.64.199.82:4145", "socks5://174.70.241.8:24398", "socks5://184.176.166.20:4145")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnHTML("span.title-list-row__row__title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	_ = c.Visit("https://www.justwatch.com/us/search?q=2001")
}