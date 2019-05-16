package view

import (
	"go-crawler/crawler/types"
	"go-crawler/crawler/frontend/model"
	common "go-crawler/crawler/zhenai/model"
	"os"
	"testing"
)

func TestS(t *testing.T)  {
	view := CreateSearchResultView("template.html")

	//template:= template.Must(
	//template.ParseFiles("view/template.html"))

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}

	page.Hits = 123

	//item := types.Item{
	//	Url:  "http://album.zhenai.com/u/108906739",
	//	Type: "zhenai",
	//	Id:   "108906739",
	//	Payload: common.Profile{
	//		Age:        34,
	//		Height:     162,
	//		Weight:     57,
	//		Income:     "3001-5000元",
	//		Gender:     "女",
	//		Name:       "安静的雪",
	//		Xinzuo:     "牧羊座",
	//		Occupation: "人事/行政",
	//		Marriage:    "离异",
	//		House:      "已够房",
	//		Hokou:      "山东菏泽",
	//		Education:  "大学本科",
	//		Car:        "未购车",
	//	},
	//}

	item := types.Item{
		Url: "http://album.zhenai.com/u/1106374945",
		Type: "zhenai",
		Id: "1106374945",
		Payload: common.Profile{
			Age: 29,
			MemberID: 1106374945,
			BasicInfo: `["未婚", "29岁", "天秤座(09.23-10.22)", "183cm", "95kg", "工作地:阿勒泰阿勒泰市", "月收入:8千-1.2万", "银行", "大学本科"]`,
			DetailInfo: `["汉族", "籍贯:山东德州", "体型:体格魁梧", "不吸烟", "社交场合会喝酒", "已购房", "已买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:一年内"]`,
			Nickname: "寻真爱",
			GenderString: "男士",
			EducationString: "大学本科",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out,page)
	if err != nil{
		panic(err)
	}
}