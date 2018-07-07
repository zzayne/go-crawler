package engine

import (
	"log"

	"github.com/zzayne/go-crawler/fetcher"
)

//Engine 初始解析入口
type Engine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

// Scheduler ...
type Scheduler interface {
	ReadyNotifier
	WorkChan() chan Request
	Run()
	Submit(Request)
}

//ReadyNotifier ...
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//Run 启动解析引擎
func (e *Engine) Run(reqs ...Request) {
	resultChan := make(chan ParseResult)

	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkChan(), resultChan, e.Scheduler)
	}

	for _, r := range reqs {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-resultChan
		for _, item := range result.Items {
			e.ItemChan <- item
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}

}

func worker(r Request) (ParseResult, error) {

	doc, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf(" fetch url error:%s\n", r.URL)
		return ParseResult{}, err
	}
	return r.ParseFunc(doc)
}

func createWorker(in chan Request, out chan ParseResult, noticer ReadyNotifier) {
	go func() {
		for {
			noticer.WorkerReady(in)
			req := <-in
			result, err := worker(req)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
