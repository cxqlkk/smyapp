package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"net/http"
	"reflect"
	"smy/controller/session"
	"smy/model"
	"smy/service"
	"smy/util"
	"strconv"
	"time"
)

type studentController struct {
}
type controllerMsg map[string]interface{}

func (stu *studentController) AddStudent(w http.ResponseWriter, r *http.Request) {
	n, err := snowflake.NewNode(1)
	birthday, _ := time.Parse("2006-01-02", r.FormValue("birthDay"))
	phoneNum := r.FormValue("phoneNum")
	if err == nil {
		jsonObj := model.Student{
			ID:          n.Generate().String(),
			StudentName: r.FormValue("studentName"),
			PhoneNum:    phoneNum,
			Ideal:r.FormValue("ideal"),
			FavoriteColor:r.FormValue("favoriteColor"),
			FavoritePerson:r.FormValue("favoritePerson"),
			Address:     r.FormValue("address"),
			BirthDay:    birthday,
			DefaultModel: model.DefaultModel{
				Status:     1,
				Created:    "管理员",
				Updated:    "管理员",
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
		}
		mess, _ := json.Marshal(new(service.StudentService).AddStudent(jsonObj))
		fmt.Fprint(w, string(mess))
	}

}

func (s *studentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	studentvo := &model.StudentVo{}
	util.FormToStruct(r, reflect.ValueOf(studentvo))
	mess := service.DefaultStudentService.UpdateStudent(*studentvo)
	bts, err := json.Marshal(mess)
	controllerCheckError(err)
	fmt.Fprintln(w,string(bts))
}

//学生端登陆
func (stu *studentController) LoginStudent(w http.ResponseWriter, r *http.Request) {

	var msg controllerMsg
	phoneNum := r.FormValue("phoneNum")
	studentName := r.FormValue("studentName")
	fmt.Println("login", phoneNum, studentName)
	student, exist := new(service.StudentService).LoginStudent(model.Student{
		StudentName: studentName,
		PhoneNum:    phoneNum,
	})
	if exist {
		sessionManger := session.DefalutSessionManger
		session := sessionManger.CreateSession()
		cookie := http.Cookie{
			Name:  sessionManger.CookieName,
			Value: session.SessionId,
			Path:  "/", MaxAge: 3600}
		http.SetCookie(w, &cookie)
		msg = controllerMsg{"success": true, "student": student}
	} else {
		msg = controllerMsg{"success": false}
	}
	js, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(js))
}

func (stu *studentController) ListStudent(w http.ResponseWriter, r *http.Request) {

	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	studentName := r.FormValue("studentName")
	fmt.Println(studentName)
	dataStore := service.DefaultStudentService.ListStudent(util.NewPage(util.WithPageNo(pageNo), util.WithPageSize(pageSize)), studentName)
	bts, error := json.Marshal(dataStore)
	controllerCheckError(error)
	fmt.Fprintln(w, string(bts))
}

func (stu *studentController) GetStudentById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	student := service.DefaultStudentService.GetStudentById(id)
	bts, err := json.Marshal(student)
	controllerCheckError(err)
	fmt.Fprintln(w, string(bts))
}
func (stu *studentController)DeleteStudentById(w http.ResponseWriter,r *http.Request){
	id:=r.FormValue("id")
	operateMsg:=service.DefaultStudentService.DeleteStudentById(id)
	bts,err:=json.Marshal(operateMsg)
	controllerCheckError(err)
	fmt.Fprintln(w,string(bts))
}

func (stu *studentController)BatchAddStudent(w http.ResponseWriter,r *http.Request){
	students:=r.FormValue("students")
	var st []map[string]string //map[interface{}]interface 也解析为空
	json.Unmarshal([]byte(students),&st) //必须指针，不然值解析不到
	fmt.Println("st",st)
	count:=0
	for _,s:=range st{
		time,_:=time.Parse("2006-01-02",s["birthDay"])
		mess:=service.DefaultStudentService.AddStudent(model.Student{StudentName:s["studentName"],PhoneNum:s["phoneNum"],BirthDay:time})
		if mess.Success{
			count++
		}
	}

	fmt.Fprintln(w,"导入成功"+strconv.Itoa(count)+"条")
}