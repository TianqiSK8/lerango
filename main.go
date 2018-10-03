package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConCurrentEngine{
		 //Scheduler: 	&scheduler.SimpleScheduler{},
		 Scheduler:		&scheduler.QueuedScheduler{},
		 WorkerCount: 	100,
		 ItemChan:		persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:		"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
	//e.Run(engine.Request{
	//	Url:		"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
