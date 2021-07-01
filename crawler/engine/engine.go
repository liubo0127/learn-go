package engine

import (
	"log"

	"github.com/liubo0127/learn-go/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetchor: error fetching url %s: %v\n", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}
