package dao

import (
	"github.com/satori/go.uuid"
	"os"
	"smy/model"
	"smy/util"
	"strings"
	"time"
)

type pictureDao struct {
}

var DefaultPictureDao = &pictureDao{}

func (pd *pictureDao) AddPicture(picture model.Picture) (int64, error) {
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("insert into tb_picture (id,studentId,fileName,realName,description,courseType,status,created,updated,createTime,updateTime) values (?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	DbError(err)
	result, err := stmt.Exec(strings.ReplaceAll(uuid.NewV4().String(), "-", ""), picture.StudentId, picture.FileName, picture.RealName, picture.Description, picture.CourseType, model.ACTIVE, "", "", time.Now(), time.Now())
	DbError(err)
	return result.RowsAffected()
}

func (pd *pictureDao) Count(studentId, courseType, createTime string) (count int) {
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select count(1) count from tb_picture inner join tb_student " +
		" on tb_student.id=tb_picture.studentId where tb_picture.status=1 and tb_student.status=1 and tb_student.id like concat('%',?,'%') and courseType like concat('%',?,'%') and DATE_FORMAT(tb_picture.createTime,'%Y-%m') like concat('%',?,'%') ")
	defer stmt.Close()
	DbError(err)
	stmt.QueryRow(studentId, courseType, createTime).Scan(&count)
	return
}
func (pd *pictureDao) ListPicture(page util.Page, studentId, courseType, createTime string) (data []model.Picture) {

	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select tb_picture.id, studentName,courseType,description,fileName,realName from tb_picture inner join tb_student " +
		" on tb_student.id=tb_picture.studentId where tb_picture.status=1 and tb_student.status=1 and tb_student.id like concat('%',?,'%') and courseType like concat('%',?,'%') and DATE_FORMAT(tb_picture.createTime,'%Y-%m') like concat('%',?,'%') limit ?,? ")
	defer stmt.Close()
	DbError(err)
	rows, err := stmt.Query(studentId, courseType, createTime, (page.PageNo-1)*page.PageNo, page.PageSize)
	for rows.Next() {
		picture := model.Picture{}
		rows.Scan(&picture.Id, &picture.StudentName, &picture.CourseType, &picture.Description, &picture.FileName, &picture.RealName)
		data = append(data, picture)
	}
	return
}

//获得时间线
func (pd *pictureDao) PictureTimeLine(studentId, courseType string) (data []model.PictureTimeLine) {
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select DATE_FORMAT(createTime,'%Y-%m') timeLine from smy.tb_picture where studentId=? and courseType=? group by timeLine")
	defer stmt.Close()
	DbError(err)
	rows, err := stmt.Query(studentId, courseType)
	DbError(err)
	for rows.Next() {
		timeLine := model.PictureTimeLine{}
		rows.Scan(&timeLine.TimeLine)
		data = append(data, timeLine)
	}
	return
}

//删除
func (pd *pictureDao) DeletePictureById(id ,realName  string)bool{
	db:=OpenDb()
	defer db.Close()
	st,err:=db.Prepare("delete from tb_picture where id =?")
	DbError(err)
	re,err:=st.Exec(id)
	DbError(err)
	rowAffect,_:=re.RowsAffected()
	if rowAffect==1{
		os.Remove(realName)
		return true
	}
	return false

}
