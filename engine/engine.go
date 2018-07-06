package engine

import (
	"fmt"
	"log"

	"github.com/zzayne/zayne-crawler/fetcher"
)

//SimpleEngine 初始解析入口
type SimpleEngine struct{}

var requestList []Request

//Run 启动解析引擎
func (e *SimpleEngine) Run(req Request) {

	requestList = append(requestList, req)
	count := 0
	for {
		var req = requestList[0]
		fmt.Printf("fetch url # %d:%s\n", count, req.URL)

		doc, err := fetcher.Fetch(req.URL)
		if err != nil {
			log.Printf(" fetch url error:%s\n", req.URL)
		}

		result, err := req.ParseFunc(doc)

		requestList = requestList[1:]

		// if err != nil {
		// 	panic(err)
		// }

		for _, m := range result {
			requestList = append(requestList, m)
		}

		if len(requestList) == 0 {
			break
		}
	}

}

func worker(req Request) {

}
