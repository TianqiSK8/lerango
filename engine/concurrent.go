package engine

import "log"

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConCurrentEngine) Run(seeds ...Request)  {
	//in := make(chan Request)
	//out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	//
	//for i := 0; i < e.WorkerCount; i++ {
	//	createWorker(in, out)
	//}
	//
	//for _, r := range seeds {
	//	e.Scheduler.Submit(r)
	//}
	//
	//itemCount := 0
	//for {
	//	result := <-out
	//	for _, item := range result.Items {
	//		log.Printf("Got item #%d: %v", itemCount, item)
	//		itemCount++
	//	}
	//
	//	for _, request := range result.Requests {
	//		e.Scheduler.Submit(request)
	//	}
	//}

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//func createWorker(in chan Request, out chan ParseResult) {
//	go func() {
//		for {
//			//tell scheduler i'm ready
//			request := <-in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//
//			out <- result
//		}
//	}()
//}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}