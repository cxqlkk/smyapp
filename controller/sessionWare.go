package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"smy/controller/session"
)

func SessionWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.BasicAuth()
		fmt.Println(r.URL)
		errMsg := make(map[string]string)
		if match, _ := regexp.MatchString("(/smy/controller/LoginStudent)|(/smy/controller/LoginTeacher)", r.RequestURI); !match { //todo
			sessionManager := session.DefalutSessionManger
			smyCookie, err := r.Cookie(sessionManager.CookieName)

			if err != nil {
				controllerCheckError(err)
				errMsg["error"] = " please Login"
			} else {
				session := sessionManager.ReadSession(smyCookie.Value)
				fmt.Println("session",session)
				if session == nil {
					errMsg["error"] = " please Login!"
				} else {
					sessionManager.SessionReset(session)
				}
			}
			if errMsg["error"] != "" {
				bts, _ := json.Marshal(errMsg)
				fmt.Fprintln(w, string(bts))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
