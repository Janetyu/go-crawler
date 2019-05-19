package data

import (
	"fmt"
	"reflect"
	"regexp"
	"gopkg.in/olivere/elastic.v5"
	"go-crawler/crawler/types"
	"go-crawler/crawler/backend/model"
	"context"
	"errors"
)

const index  = "knowledge_base"

type Item struct {
	Url string `json:"Url"`
	Type string `json:"Type"`
	Id string `json:"Id"`
	Payload interface{} `json:"Payload"`
}

type Article struct {
	Title string `json:"Title"`
	Author string `json:"Author"`
	Abstract string `json:"Abstract"`
	Label string `json:"Label"`
	Type string `json:"Type"`
	Date string `json:"Date"`
}

type SearchResult struct {
	Hits int64 `json:"hits"`
	Start int `json:"start"`
	Items []interface{} `json:"items"`
	//Items []types.Item
	Query string `json:"query"`
	PrevFrom int `json:"prev_from"`
	NextFrom int `json:"next_from"`
}

func GetSearchResult(q string, from int) (SearchResult, error) {
	var result SearchResult
	result.Query = q
	resp, err := model.Client.
		Search("knowledge_base").
		//对q 进行正则替换，只在给elasticSearch时rewrite
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	//己有的包装
	result.Hits = resp.TotalHits()
	result.Start = from
	//用到了反射
	//for _, v := range resp.Each(reflect.TypeOf(types.Item{})) {
	//	item := v.(types.Item)
	//	var Profile common.Profile
	//	item.Payload = mapstructure.Decode(item.Payload, &Profile)
	//	v = item
	//}
	result.Items = resp.Each(reflect.TypeOf(types.Item{}))
	fmt.Printf("the get data result :%v \n",result)


	//支持分页
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

//前端搜索没有输入Payload.Age<30，而是输入的Age<30,因此为了保证能在elastic中正常使用，重新添加上Payload.
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`) //([A-Z][a-z]*)表示Height: Age:
	//$1 代表（）中的部分
	return re.ReplaceAllString(q, "Payload.$1:")
}


func SaveOrUpdate(index string, item Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := model.Client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return err
}

func Delete(index string, item Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	deleteService := model.Client.Delete().
		Index(index).
		Type(item.Type).
		Id(item.Id)

	if item.Id != "" {
		deleteService.Id(item.Id)
	}

	_, err := deleteService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return err
}
