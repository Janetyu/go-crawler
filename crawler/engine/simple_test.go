package engine_test

import (
	"testing"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/types"
)

func TestSimpleEngine_Run(t *testing.T) {
	engine.SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
