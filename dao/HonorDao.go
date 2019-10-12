package dao

import (
	"fmt"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"smy/model"
	"smy/util"
	"strings"
	"time"
)

type honorDao struct {
}

var DefaultHonorDao = &honorDao{}

func (h *honorDao) AddHonor(honor model.Honor) (bool, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("insert into tb_honor (id,studentId,honorName,acquisitionTime,status,created,updated,createTime,updateTime) values (?," +
		"?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		logger.Info("daoError", zap.String("msg", err.Error()))
		return false, err
	}
	result, err := stmt.Exec(strings.ReplaceAll(uuid.NewV4().String(), "-", ""), honor.StudentId, honor.HonorName, honor.AcquisitionTime, model.ACTIVE, "", "", time.Now(), time.Now())
	if err != nil {
		logger.Info("daoError", zap.String("msg", err.Error()))
		return false, err
	}
	rowAffected, err := result.RowsAffected()
	return rowAffected == 1, err

}
func (h *honorDao) ListHonorByStudentId(studentId string) (datas []model.Honor) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select id,studentId,honorName,acquisitionTime from tb_honor where studentId =? and status=1 order by acquisitionTime desc")
	defer stmt.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	rows, err := stmt.Query(studentId)
	if err != nil {
		logger.Error(err.Error())
	}

	for rows.Next() {
		var id, stuId, honorName string
		var acquisitionTime time.Time
		rows.Scan(&id, &stuId, &honorName, &acquisitionTime)
		datas = append(datas, model.Honor{
			Id: id, StudentId: stuId, HonorName: honorName, AcquisitionTime: acquisitionTime,
		})
		fmt.Println(acquisitionTime)
	}
	return datas
}

func (h *honorDao) Count(studentName, honorName string) int {
	db := OpenDb()
	defer db.Close()
	stmt, err := db.Prepare("select count(1) count from tb_honor" +
		" inner join tb_student on tb_honor.studentId = tb_student.id " +
		"where tb_honor.status=1 and  tb_student.status=1 and studentName like concat('%',?,'%') and honorName like concat('%',?,'%')")
	defer stmt.Close()
	DbError(err)
	var count int
	stmt.QueryRow(studentName, honorName).Scan(&count)
	return count
}

func (h *honorDao) ListHonor(page *util.Page,studentName,honorName string)(datas []model.Honor) {
	db:=OpenDb()
	defer db.Close()
	stmt,err:=db.Prepare("select  tb_student.studentName,tb_honor.honorName,tb_honor.acquisitionTime from tb_honor"+
	" inner join tb_student on tb_honor.studentId = tb_student.id"+
	" where tb_honor.status=1 and  tb_student.status=1 and studentName like concat('%',?,'%') and honorName like concat('%',?,'%') limit ?,?")
	defer stmt.Close()
	DbError(err)
	rows,err:=stmt.Query(studentName,honorName,(page.PageNo-1)*page.PageSize,page.PageSize)
	DbError(err)
	for rows.Next(){
		ho:=model.Honor{}
		rows.Scan(&ho.StudentName,&ho.HonorName,&ho.AcquisitionTime)
		datas=append(datas, ho)
	}
	return datas
}

func (h *honorDao) DeleteHonor(honor model.Honor) {

}

func (h *honorDao) updateHonor(honor model.Honor) {

}
