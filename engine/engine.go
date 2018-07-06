package engine

import (
	"fmt"

	"github.com/zzayne/zayne-crawler/fetcher"
)

//SimpleEngine 初始解析入口
type SimpleEngine struct{}

//Run 启动解析引擎
func (e *SimpleEngine) Run(req Request) {
	var requestList []Request

	requestList = append(requestList, req)
	count := 0
	for {
		var req = requestList[0]
		fmt.Printf("fetch url # %d:%s\n", count, req.URL)

		doc, err := fetcher.Fetch(req.URL)
		if err != nil {
			panic(err)
		}

		result, err := req.ParseFunc(doc)

		requestList = requestList[1:]

		if err != nil {
			panic(err)
		}

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
