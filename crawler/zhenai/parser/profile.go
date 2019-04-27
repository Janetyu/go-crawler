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

func ParseProfile(contents []byte) types.ParseResult {
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
		Items: []interface{}{profile},
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
