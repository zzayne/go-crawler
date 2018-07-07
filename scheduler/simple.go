package scheduler

import "github.com/zzayne/go-crawler/engine"

//SimpleScheduler 每个worker公用一个chanel来接受数据，其实没有进行调度
type SimpleScheduler struct {
	workChan chan engine.Request
}

//WorkChan ...
func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workChan
}

//Run 初始化
func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

//Submit ...
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()

}
