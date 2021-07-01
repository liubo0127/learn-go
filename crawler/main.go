package main

import (
	"github.com/liubo0127/learn-go/crawler/engine"
	"github.com/liubo0127/learn-go/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

//type User struct {
//	Name     string
//	Age      string
//	Height   string
//	HasHouse bool
//	HasCar   bool
//}
