package controller

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"smy/config"
	"smy/model"
	"smy/service"
	"smy/util"
	"strconv"
	"strings"
	"time"
)

type pictureController struct{}

func (p *pictureController) AddPicture(w http.ResponseWriter, r *http.Request) {
	picture := model.Picture{
		StudentId:   r.FormValue("studentId"),
		Description: r.FormValue("description"),
		CourseType:  r.FormValue("courseType"),
		DefaultModel: model.DefaultModel{
			Status:     model.ACTIVE,
			Created:    "管理员",
			Updated:    "管理员",
			CreateTime: time.Now(),
		},
	}

	bts, error := json.Marshal(service.DefaultPictureService.AddPicture(picture))
	controllerCheckError(error)
	fmt.Fprintln(w, string(bts))
}

//上传图片
func (p *pictureController) UploadPicture(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20);
	//获取上传的文件组
	files := r.MultipartForm.File["file"];
	len := len(files);
	fmt.Println("文件数", len)
	studentId := r.FormValue("studentId")
	courseType := r.FormValue("courseType")
	description := r.FormValue("description")
	for _, file := range files {
		lastPoint := strings.LastIndex(file.Filename, ".")
		serverFileName := strings.ReplaceAll(uuid.NewV4().String(), "-", "") + file.Filename[lastPoint:]
		serverFile, err := os.Create(config.UpLoadPath + "/" + serverFileName)
		controllerCheckError(err)
		of, _ := file.Open()
		defer of.Close()
		io.Copy(serverFile, of)
		service.DefaultPictureService.AddPicture(model.Picture{
			StudentId:   studentId,
			CourseType:  courseType,
			Description: description,
			RealName:    serverFileName,
			FileName:    file.Filename,
		})
	}
	bts, _ := json.Marshal(util.OperateMessage{Success: true, Message: "添加成功"})
	fmt.Fprintln(w, string(bts))
}
//学生的图片分页查询
func (pc *pictureController)ListPicture(w http.ResponseWriter,r *http.Request){
	pageNo,_:=strconv.Atoi(r.FormValue("pageNo"))
	pageSize,_:=strconv.Atoi(r.FormValue("pageSize"))
	page:=util.NewPage(util.WithPageNo(pageNo),util.WithPageSize(pageSize))
	studentId:=r.FormValue("studentId")
	courseType:=r.FormValue("courseType")
	createTime:=r.FormValue("createTime")
	dataStore:=service.DefaultPictureService.ListPicture(*page,studentId,courseType,createTime)
	bts,err:=json.Marshal(dataStore)
	controllerCheckError(err)
	fmt.Fprintln(w,string(bts))

}
//picture timeLine
func(pc *pictureController)PictureTimeLine(w http.ResponseWriter,r *http.Request){
	studentId:=r.FormValue("studentId")
	courseType:=r.FormValue("courseType")
	data:=service.DefaultPictureService.PictureTimeLine(studentId,courseType)
	bts,err:=json.Marshal(data)
	controllerCheckError(err)
	fmt.Fprintln(w,string(bts))
}

//删除
func (pc *pictureController)DeletePictureById(w http.ResponseWriter,r *http.Request){
	id:=r.FormValue("id")
	realName:=r.FormValue("realName")
	mes:=service.DefaultPictureService.DeletePictureById(id,realName)
	bts,_:=json.Marshal(mes)
	fmt.Fprintln(w,string(bts))

}