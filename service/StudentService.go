package service

import (
	"smy/dao"
	"smy/model"
	"smy/util"
	"strconv"
)

type StudentService struct {
}

var DefaultStudentService = &StudentService{}

func (s *StudentService) AddStudent(stu model.Student) util.OperateMessage {
	row, err := new(dao.StudentDao).AddStudent(stu)
	if err != nil || row == 0 {
		return util.OperateMessage{
			Success: false,
			Message: "操作失败",
		}
	}
	return util.OperateMessage{
		Success: true,
		Message: "操作成功",
	}
}

func (s *StudentService) LoginStudent(stu model.Student) (model.Student, bool) {
	return new(dao.StudentDao).LoginStudent(stu)
}

//列表查询
func (s *StudentService) ListStudent(page *util.Page, studentName string) (dataStore util.DataStore) {
	total := dao.DefaultStudentDao.Count(studentName)
	data := dao.DefaultStudentDao.ListStudent(page, studentName)
	for _, da := range data { //不能将[]student 赋值给 []interface
		dataStore.Datas = append(dataStore.Datas, da)
	}
	dataStore.Total=total


	return dataStore
}

func (s *StudentService)GetStudentById(id string)model.Student{
	return dao.DefaultStudentDao.GetStudentById(id);
}
func (s *StudentService)UpdateStudent(vo model.StudentVo)( operateMessage util.OperateMessage){
	row,error:=dao.DefaultStudentDao.UpdateStudent(vo)
	if error==nil{
		dao.DbError(error)
	return util.OperateMessage{
			Success:true,
			Message:"修改成功",
		}
	}
	return util.OperateMessage{
		Success:true,
		Message:"修改成功"+strconv.Itoa(int(row))+"条",
	}
}
func (s *StudentService)DeleteStudentById(id string)util.OperateMessage{
	row,_:=dao.DefaultStudentDao.DeleteStudentById(id)
	if row==1{
		return util.OperateMessage{
			Success:true,
			Message:"删除成功",
		}
	}
	return util.OperateMessage{
		Success:false,
		Message:"删除失败",
	}
}