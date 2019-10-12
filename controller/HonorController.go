package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"smy/model"
	"smy/service"
	"smy/util"
	"strconv"
	"time"
)

type honorController struct {
}

func (h *honorController) AddHonor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addHonor")
	id := r.FormValue("id")
	studentId := r.FormValue("studentId")
	honorName := r.FormValue("honorName")
	acquisitionTime, _ := time.Parse("2006-01-02", r.FormValue("acquisitionTime"))
	honor := model.Honor{
		Id:              id,
		StudentId:       studentId,
		HonorName:       honorName,
		AcquisitionTime: acquisitionTime,
		DefaultModel: model.DefaultModel{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			Created:    "",
			Updated:    "",
			Status:     model.ACTIVE,
		},
	}
	byteMsg, err := json.Marshal(service.DefaultHonorService.AddHonor(honor))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(byteMsg))
}

func (h *honorController) ListHonor(w http.ResponseWriter, r *http.Request) {
	pageSize ,_:=strconv.Atoi(r.FormValue("pageSize"))
	pageNo,_:=strconv.Atoi(r.FormValue("pageNo"))
	studentName:=r.FormValue("studentName")
	honorName:=r.FormValue("honorName")
	page:=util.NewPage(util.WithPageSize(pageSize),util.WithPageNo(pageNo))
	dataStore:=service.DefaultHonorService.ListHonor(page,studentName,honorName)
	bts,err:=json.Marshal(dataStore)
	controllerCheckError(err)
	fmt.Fprintln(w,string(bts))

}
func (h *honorController) ListHonorByStudentId(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.FormValue("studentId"),"sdf")
	datas := service.DefaultHonorService.ListHonorByStudentId(r.FormValue("studentId"))
	bts, err := json.Marshal(datas)
	controllerCheckError(err)
	fmt.Fprintln(w, string(bts))
}

func (h *honorController) DeleteHonor(w http.ResponseWriter, r *http.Request) {

}

func (h *honorController) UpdateHonor(w http.ResponseWriter, r *http.Request) {

}
