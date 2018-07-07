package engine

import (
	"log"

	"github.com/zzayne/zayne-crawler/fetcher"
	"github.com/zzayne/zayne-crawler/persist"
)

//SimpleEngine 初始解析入口
type SimpleEngine struct{}

//Run 启动解析引擎
func (e *SimpleEngine) Run(reqs ...Request) {
	var requests []Request
	for _, r := range reqs {
		requests = append(requests, r)
	}
	count := 0
	for {
		if len(requests) > 0 {
			var req = requests[0]
			requests = requests[1:]

			result, err := worker(req)

			if err != nil {
				continue
			}

			requests = append(requests, result.Requests...)
			for _, item := range result.Items {
				//log.Printf("got item # %d:%v\n", count, item)
				persist.ItemSave(item)
				count++
			}
		}
	}
}

func worker(r Request) (ParseResult, error) {
	doc, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf(" fetch url error:%s\n", r.URL)
	}
	return r.ParseFunc(doc)
}
