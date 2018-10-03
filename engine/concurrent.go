package engine

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	//通过chan存储爬取下来的item
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConCurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//把item送出去
			go func() {
				e.ItemChan <- item
			}()
		}
		//for i:=0;i<len(result.Items);i++  {
		//	it := &result.Items[i]
		//	go func() {
		//		e.ItemChan <- *it
		//	}()
		//}

		for _, request := range result.Requests {
			//TO DO 去重方法
			//if isDuplicate(request.Url){
			//	continue
			//}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}