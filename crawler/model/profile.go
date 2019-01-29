package model

type Profile struct {
	MemberID int `json:"memberID"`
	BasicInfo string `json:"basicInfo"`
	DetailInfo string `json:"detailInfo"`
	Age int `json:"age"`
	Nickname string `json:"nickname"`
	GenderString string `json:"genderString"`
	EducationString string `json:"educationString"`
}