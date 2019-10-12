package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smy/controller/session"
	"smy/service"
)

type teacherController struct {
}

func (tc *teacherController) LoginTeacher(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	password := r.FormValue("password")
	var mess controllerMsg
	teacher := service.DefaultTeacherService.QueryTeacher(account, password)
	if teacher.Account != "" {
		mess = controllerMsg{
			"success": true,
			"teacher": teacher,
		}
		sessionManager := session.DefalutSessionManger
		ses := sessionManager.CreateSession()
		cookie := http.Cookie{
			Name:  sessionManager.CookieName,
			Value: ses.SessionId,
			Path:  "/", MaxAge: 3600}
		http.SetCookie(w, &cookie)
	} else {
		mess = controllerMsg{
			"success": false,
		}
	}
	bts, err := json.Marshal(mess)
	controllerCheckError(err)
	fmt.Fprintln(w, string(bts))
}
