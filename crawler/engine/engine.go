package engine

import (
	"log"
	"go-crawler/crawler/fetcher"
	"time"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}


	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		// parseResult.Requests... 意思是把[]Requests 拆开逐个添加
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

		time.Sleep(2 * time.Second)
	}
}
