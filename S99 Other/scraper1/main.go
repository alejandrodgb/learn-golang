package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("Ping\n"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	// Verify the parameter url exists
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("Missing URL argument")
		return
	}
	log.Println("Visiting", URL)

	//Create a new collector which will be in charge of collect the data from HTML
	c := colly.NewCollector()

	// Slices to store the data
	var response []string

	//onHTML function allows the collector to use a callback function when the specific HTML tag is reached
	//in this case whenever our collector finds an
	//anchor tag with href it will call the anonymous function
	// specified below which will get the info from the href and append it to our slice
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			response = append(response, link)
		}
	})

	//Command to visit the website
	c.Visit(URL)

	// parse our response slice into JSON format
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	// Add some header and write the body for our endpoint
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	addr := ":7171"

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/search", getData)

	log.Println("Listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
