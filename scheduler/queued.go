package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (qs *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (qs *QueuedScheduler) Submit(request engine.Request) {
	qs.requestChan <- request
}

func (qs *QueuedScheduler) WorkerReady(w chan engine.Request) {
	qs.workerChan <- w
}

func (qs *QueuedScheduler) Run() {
	qs.workerChan  = make(chan chan engine.Request)
	qs.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ  []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker  = workerQ[0]
			}
			select {
			case r := <-qs.requestChan:
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ  = workerQ[1:]
			}
		}
	}()
}

