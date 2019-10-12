package model

type Teacher struct {
	Id string `json:"id"`
	TeacherName string `json:"teacherName"`
	Account string `json:"account"`
	Password string `json:"password"`
	DefaultModel
}
