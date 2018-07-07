package engine

import (
	"log"

	"github.com/zzayne/go-crawler/fetcher"
	"github.com/zzayne/go-crawler/persist"
)

//SimpleEngine 初始解析入口
type SimpleEngine struct{}

//Run 启动解析引擎
func (e *SimpleEngine) Run(reqs ...Request) {
	workChan := make(chan Request)
	resultChan := make(chan ParseResult)

	for i := 0; i < 10; i++ {
		createWorker(workChan, resultChan)
	}

	count := 0

	for _, r := range reqs {
		go func() {
			workChan <- r
		}()
	}

	for {
		result := <-resultChan

		for _, item := range result.Items {
			log.Printf("got item # %d:%v\n", count, item)
			persist.ItemSave(item)
			count++
		}

		for _, r := range result.Requests {
			go func() {
				workChan <- r
			}()
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			req := <-in

			result, err := worker(req)
			if err != nil {
				continue
			}
			log.Println(result)
			out <- result
		}
	}()
}
