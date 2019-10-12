package service

import (
	"smy/dao"
	"smy/model"
)

type teacherService struct {

}

var DefaultTeacherService =&teacherService{}

func (ts *teacherService)QueryTeacher(account,password string)model.Teacher{
	return dao.DefaultTeacherDao.QueryTeacher(account,password)
}