package engine

import (
	"go-crawler/crawler/types"
	"go-crawler/crawler/fetcher"
	"log"
)

func worker(r types.Request) (types.ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return types.ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}

