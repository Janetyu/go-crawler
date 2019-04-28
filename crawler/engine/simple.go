package engine

import (
	"log"
	"time"
	"go-crawler/crawler/fetcher"
	"go-crawler/crawler/types"
)

type SimpleEngine struct{}

func (e SimpleEngine)Run(seeds ...types.Request) {
	var requests []types.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 逐个 request 进行处理
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// parseResult.Requests... 意思是把[]Requests 拆开逐个添加
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

		time.Sleep(3 * time.Second)
	}
}

func worker(r types.Request) (types.ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return types.ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
