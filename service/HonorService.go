package service

import (
	"smy/dao"
	"smy/model"
	"smy/util"
)

type honorService struct {
}

var DefaultHonorService = &honorService{}
/**
@description 添加荣誉
*/
func (hs *honorService) AddHonor(honor model.Honor) util.OperateMessage {
	success, _ := dao.DefaultHonorDao.AddHonor(honor)
	if !success {
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

/**
@description 获得指定学生的荣誉列表
*/
func (hs *honorService) ListHonorByStudentId(studentId string) []model.Honor {

	return dao.DefaultHonorDao.ListHonorByStudentId(studentId)
}
/**
分页查询
 */
func (hs *honorService)ListHonor(page *util.Page,studentName,honorName string)(datastore util.DataStore){
	count:=dao.DefaultHonorDao.Count(studentName,honorName)
	datas:=dao.DefaultHonorDao.ListHonor(page ,studentName,honorName)
	datastore.Total=count
	for _,da:=range datas{
		datastore.Datas=append(datastore.Datas, da)
	}
	return datastore
}
