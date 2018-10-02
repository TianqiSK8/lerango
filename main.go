package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:		"http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//})

	e := engine.ConCurrentEngine{
		 Scheduler: 	&scheduler.SimpleScheduler{},
		 WorkerCount: 	100,
	}

	e.Run(engine.Request{
		Url:		"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
