package scheduler

import "github.com/zzayne/go-crawler/engine"

//QueueScheduler 为每个worker单独分配一个chanel，由调度器来按照队列的方式，对worker用到的chanel和request进行对应分配
type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

//WorkerReady ...
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//Submit ...
func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//WorkChan ...
func (s *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

//Run ...
func (s *QueueScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeR engine.Request
			var activeW chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeR = requestQ[0]
				activeW = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeW <- activeR:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}

		}

	}()

}
