package service

import (
	"smy/dao"
	"smy/model"
	"smy/util"
)

type pictureService struct {
}

var DefaultPictureService = &pictureService{}

func (ps *pictureService) AddPicture(picture model.Picture) util.OperateMessage {
	rows, _ := dao.DefaultPictureDao.AddPicture(picture)
	if rows == 1 {
		return util.OperateMessage{
			Success: true,
			Message: "添加成功",
		}
	}
	return util.OperateMessage{
		Success: false,
		Message: "添加失败",
	}
}

//
func (ps *pictureService) ListPicture(page util.Page, studentId, courseType, createTime string) (dataStore util.DataStore) {
	dataStore.Total = dao.DefaultPictureDao.Count(studentId, courseType, createTime)
	data := dao.DefaultPictureDao.ListPicture(page, studentId, courseType, createTime)
	for _, d := range data {
		dataStore.Datas = append(dataStore.Datas, d)
	}
	return
}
func (ps *pictureService) PictureTimeLine(studentId, courseType string) []model.PictureTimeLine {
	return dao.DefaultPictureDao.PictureTimeLine(studentId, courseType)
}

func (ps *pictureService) DeletePictureById(id, realName string) util.OperateMessage {
	b := dao.DefaultPictureDao.DeletePictureById(id, realName)
	if b {
		return util.OperateMessage{
			Success: true,
			Message: "删除成功",
		}
	}

	return util.OperateMessage{
		Success: true,
		Message: "删除成功",
	}
}
