package model

import "encoding/json"

type Profile struct {
	MemberID int `json:"memberID"`
	BasicInfo string `json:"basicInfo"`
	DetailInfo string `json:"detailInfo"`
	Age int `json:"age"`
	Nickname string `json:"nickname"`
	GenderString string `json:"genderString"`
	EducationString string `json:"educationString"`
}

func FromJson2Obj(o interface{}) (Profile, error)  {
	var profile Profile
	// format o to string
	s, err := json.Marshal(o)
	if err !=nil {
		return profile, err
	}

	// format string to obj
	json.Unmarshal(s, &profile)
	return profile,err
}