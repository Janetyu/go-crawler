package parser

import (
	"regexp"
	"go-crawler/crawler/zhenai/model"
	"strconv"
	"go-crawler/crawler/types"
)

var (
//getServerDataUrl  = `http://album.zhenai.com/api/profile/getObjectProfile.do?_=1548668275798&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2Fb39a9ff7-3dd0-4c81-a301-ce6934c221a4%2F0%2F0%2F1875468050&objectID=`
memberIdRe  = regexp.MustCompile(`"memberID":([^,]+)`)
ageRe = regexp.MustCompile(`"age":([^,]+),`)
genderRe = regexp.MustCompile(`"genderString":"([^"]+)",`)
educatedRe = regexp.MustCompile(`"educationString":"([^"]+)",`)
nicknameRe = regexp.MustCompile(`"nickname":"([^"]+)",`)
basicInfoRe = regexp.MustCompile(`"basicInfo":([^\]]+])`)
detailInfoRe = regexp.MustCompile(`"detailInfo":([^\]]+])`)
)

func ParseProfile(contents []byte, url string) types.ParseResult {
	//resp, err := http.Get(getServerDataUrl + objectId)
	//defer resp.Body.Close()
	profile := model.Profile{}

	//fmt.Print(string(contents))

	age, err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil {
		profile.Age = age
	}

	memberId,err := strconv.Atoi(extractString(contents,memberIdRe))
	if err == nil {
		profile.MemberID = memberId
	}

	gender := extractString(contents,genderRe)
	profile.GenderString = gender

	educated := extractString(contents,educatedRe)
	profile.EducationString = educated

	nickname := extractString(contents,nicknameRe)
	profile.Nickname = nickname

	basicInfo := extractString(contents,basicInfoRe)
	profile.BasicInfo = basicInfo

	detailInfo := extractString(contents,detailInfoRe)
	profile.DetailInfo = detailInfo

	result := types.ParseResult{
		Items: []types.Item{
			{
				Url: url,
				Id: strconv.Itoa(memberId),
				Type: "zhenai",
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	// 查找出第一个最匹配的结果
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

//const ageRe = `<td><span class="label">年龄：</span>(\d+)岁</td>`
//var ageRex = regexp.MustCompile(ageRe)
//var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
//var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`)
//var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
//var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
//var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^>]+)</span></td>`)
//
//const marriageRe = `<td><span class="label">婚况：</span>([^>]+)</td>`
//var marriageRex = regexp.MustCompile(marriageRe)
//
//var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`)
//var OccupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
//var HokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^>]+)</td>`)
//var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
//var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)
//
////猜你喜欢的,提取url, name
//var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<])`)
////提取id
//var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
//
//
//func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
//	profile := model.Profile{}
//	//re := regexp.MustCompile(ageRe)
//	//match := ageRex.FindSubmatch(contents)
//	//
//	//if match != nil {
//	//	age, e := strconv.Atoi(string(match[1]))
//	//	if e != nil{
//	//		profile.Age = age
//	//	}
//	//}
//	profile.Name = name
//
//	age, e := strconv.Atoi(extractString(contents, ageRex))
//	if e == nil {
//		profile.Age = age
//	}
//	height, e := strconv.Atoi(extractString(contents, heightRe))
//	if e == nil {
//		profile.Height = height
//	}
//	weight, e := strconv.Atoi(extractString(contents, weightRe))
//	if e == nil {
//		profile.Weight = weight
//	}
//
//	//re = regexp.MustCompile(marriageRe)
//	//match = marriageRex.FindSubmatch(contents)
//	//
//	//if match != nil {
//	//	profile.Marriage = string(match[1])
//	//}
//	profile.Income = extractString(contents,incomeRe)
//	profile.Gender = extractString(contents,genderRe)
//	profile.Car = extractString(contents,CarRe)
//	profile.Xinzuo = extractString(contents,xinzuoRe)
//	profile.Education = extractString(contents,educationRe)
//	profile.Occupation = extractString(contents,OccupationRe)
//	profile.Hokou = extractString(contents,HokouRe)
//	profile.House = extractString(contents,HouseRe)
//	profile.Marriage = extractString(contents,marriageRex)
//
//	result := engine.ParseResult{
//		Items:[]engine.Item{
//			{
//				Url:url,
//				Type:"zhenai",
//				Id:extractString([]byte(url),idUrlRe),
//				Payload:profile,
//			},
//		},
//	}
//	//猜你喜欢的人
//	matches := guessRe.FindAllSubmatch(contents, -1)
//	for _, m :=range matches {
//		url :=string(m[1])
//		name :=string(m[2])
//		result.Requests = append(result.Requests,
//			engine.Request{
//				Url:url,
//				//ParserFunc: ProfileParser(name),
//				Parser:NewProfileParser(name),
//			})
//	}
//
//	return result
//}
//
//func extractString(contents []byte, re *regexp.Regexp) string  {
//	match := re.FindSubmatch(contents)
//	if len(match) >=2 {
//		return string(match[1])
//	} else {
//		return ""
//	}
//
//}
//
//type ProfileParser struct {
//	userName string
//}
//
//func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
//	return ParseProfile(contents,url,p.userName)
//
//}
//
//func (p *ProfileParser) Serialize() (name string, args interface{}) {
//	//return "ProfileParser", p.userName
//	return config.ParseProfile, p.userName //
//}
//
//func NewProfileParser(name string) *ProfileParser {
//	return &ProfileParser{
//		userName:name,
//	}
//}
