package dao

import (
	"fmt"
	"smy/model"
)

type teacherDao struct {

}

var DefaultTeacherDao =&teacherDao{}

func (td *teacherDao)QueryTeacher(account,password string)( model.Teacher){
	teacher :=model.Teacher{}
	db:=OpenDb()
	defer db.Close()
	stmt,err:=db.Prepare("select id,teacherName,account from tb_teacher where status =1 and account =? and password=?")
	defer stmt.Close()
	DbError(err)

	fmt.Println(account,password)
	err=stmt.QueryRow(account,password).Scan(&teacher.Id,&teacher.TeacherName,&teacher.Account)
	DbError(err)
	return teacher
}