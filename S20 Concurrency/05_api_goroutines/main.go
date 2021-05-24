package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"sync"
)

type stats struct {
	min int
	max int
}

// URLResponse is a struct that contains the URL and response received
type URLResponse struct {
	URL      string
	Response string
}

// Array that will hold the responses
var response []URLResponse

// Array that will hold the goroutine numbers
var goRoutines []int

// URLs to review
var urls = []string{
	"https://google.com",
	"https://youtube.com",
	"https://facebook.com",
	"https://baidu.com",
	"https://wikipedia.org",
	"https://twitter.com",
	"https://yahoo.com",
	"https://pornhub.com",
	"https://instagram.com",
	"https://xvideos.com",
	"https://yandex.ru",
	"https://ampproject.org",
	"https://xnxx.com",
	"https://amazon.com",
	"https://live.com",
	"https://vk.com",
	"https://netflix.com",
	"https://qq.com",
	"https://whatsapp.com",
	"https://mail.ru",
	"https://reddit.com",
	"https://yahoo.co.jp",
	"https://google.com.br",
	"https://bing.com",
	"https://ok.ru",
	"https://xhamster.com",
	"https://sogou.com",
	"https://ebay.com",
	"https://bit.ly",
	"https://twitch.tv",
	"https://linkedin.com",
	"https://samsung.com",
	"https://sm.cn",
	"https://msn.com",
	"https://office.com",
	"https://globo.com",
	"https://taobao.com",
	"https://pinterest.com",
	"https://google.de",
	"https://microsoft.com",
	"https://accuweather.com",
	"https://naver.com",
	"https://aliexpress.com",
	"https://fandom.com",
	"https://quora.com",
	"https://github.com",
	"https://imdb.com",
	"https://uol.com.br",
	"https://docomo.ne.jp",
	"https://youporn.com",
	"https://bbc.co.uk",
	"https://microsoftonline.com",
	"https://paypal.com",
	"https://google.fr",
	"https://yidianzixun.com",
	"https://wordpress.com",
	"https://news.google.com",
	"https://sohu.com",
	"https://duckduckgo.com",
	"https://google.co.uk",
	"https://10086.cn",
	"https://iqiyi.com",
	"https://booking.com",
	"https://amazon.co.jp",
	"https://cricbuzz.com",
	"https://taboola.com",
	"https://amazon.de",
	"https://cnn.com",
	"https://jd.com",
	"https://apple.com",
	"https://google.it",
	"https://bilibili.com",
	"https://google.co.jp",
	"https://livejasmin.com",
	"https://tmall.com",
	"https://news.yahoo.co.jp",
	"https://youtu.be",
	"https://tribunnews.com",
	"https://amazon.co.uk",
	"https://chaturbate.com",
	"https://google.co.in",
	"https://craigslist.org",
	"https://imgur.com",
	"https://bbc.com",
	"https://fc2.com",
	"https://tsyndicate.com",
	"https://redtube.com",
	"https://tumblr.com",
	"https://foxnews.com",
	"https://rakuten.co.jp",
	"https://google.es",
	"https://outbrain.com",
	"https://discordapp.com",
	"https://amazon.in",
	"https://crptgate.com",
	"https://weather.com",
	"https://toutiao.com",
	"https://youku.com",
	"https://adobe.com",
	"https://news.yandex.ru",
}

func minMax(v []int) (r stats) {
	min := v[0]
	max := v[0]
	// Find minimum value
	for i, e := range v {
		if i == 0 || e < min {
			min = e
		}
	}
	// Find maximum falue
	for _, e := range v {
		if e > max {
			max = e
		}
	}
	return stats{
		min: min,
		max: max,
	}
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	log.Println("Processing url: ", url)
	rS := "ERROR"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	} else {
		rS = resp.Status
	}

	response = append(response, URLResponse{url, rS})

	wg.Done()

	return url, err
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Println("Request received")

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
		goRoutines = append(goRoutines, runtime.NumGoroutine())
	}

	wg.Wait()

	b, err := json.Marshal(response)

	if err != nil {
		log.Println("Failed to serialize response:", err)
		return
	}

	// Add some header and write the body for endpoint response
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

	log.Println("Min Goroutines:", minMax(goRoutines).min)
	log.Println("Max Goroutines:", minMax(goRoutines).max)
	log.Println("Total number of sites", len(goRoutines))

	log.Println("Program finished")
}

func handleRequests() {
	log.Println("Program started")
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
