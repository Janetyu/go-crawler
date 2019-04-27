package engine_test

import (
	"testing"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler/engine"
)

func TestSimpleEngine_Run(t *testing.T) {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
