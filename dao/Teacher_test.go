package dao

import (
	"fmt"
	"testing"
)

func TestTeacherDao_QueryTeacher(t *testing.T) {
te:=	DefaultTeacherDao.QueryTeacher("","")
fmt.Println(te==nil)
}
