package engine

import (
	"log"

	"github.com/zzayne/zayne-crawler/fetcher"
)

//SimpleEngine 初始解析入口
type SimpleEngine struct{}

var requests []Request

//Run 启动解析引擎
func (e *SimpleEngine) Run(reqs ...Request) {
	for _, r := range reqs {
		requests = append(requests, r)
	}
	count := 0
	for {
		if len(requests) > 0 {
			var req = requests[0]
			requests = requests[1:]

			doc, err := fetcher.Fetch(req.URL)
			if err != nil {
				log.Printf(" fetch url error:%s\n", req.URL)
			}

			result, err := req.ParseFunc(doc)

			requests = append(requests, result.Requests...)

			for _, item := range result.Items {
				log.Printf("got item # %d:%v\n", count, item)
				count++
			}
		}
	}
}
