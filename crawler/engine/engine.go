package engine

import (
	"crawler/fetcher"
	"log"
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
			log.Printf("Fetcher: error fetching url %s : %v", r.Url, err)
			continue
		}
		ParseResult := r.Parserfunc(body)
		// 三个点代表把这个slice的每个元素都append进去
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
