package model

type Picture struct {
	Id string `json:"id"`
	StudentId string`json:"studentId"`
	StudentName string `json:"studentName"`
	Description string `json:"description"`
	CourseType string `json:"courseType"`
	FileName string `json:"fileName"`
	RealName string `json:"realName"`
	DefaultModel
}

type PictureTimeLine struct {
	TimeLine string `json:"timeLine"`
}