package session

import (
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type session struct {
	SessionId string
	element map[interface{}]interface{}
	createdTime time.Time
	mux sync.RWMutex
	timeMux sync.RWMutex
}

type sessionManager struct {
	CookieName string
	sessions map[string]*session
	activeTime int
	mux sync.RWMutex

}
var DefalutSessionManger =&sessionManager{CookieName:"sessionId",sessions:make(map[string]*session),activeTime:3600*24}
// 初始化 全局session管理器
func NewSessionManger(cookieName string)*sessionManager{
	return &sessionManager{CookieName:cookieName}
}

func ( m *sessionManager)ReadSession(sessionId string)( s *session){
	defer m.mux.RUnlock()
	m.mux.RLock()
	s=m.sessions[sessionId]
	return
}
func (m *sessionManager)CreateSession()*session{
	sess:=&session{
		SessionId:sessionId(),
		createdTime:time.Now(),
	}
	m.mux.Lock()

	m.sessions[sess.SessionId]=sess
	m.mux.Unlock()
	return sess
}

func (m *sessionManager)SessionReset(ses *session){
	ses.createdTime=time.Now()
}

func (m *sessionManager)SessionGc(){
	for{
		select {
			case <-time.Tick(time.Second*60):
				for k,v:=range m.sessions{
					if int(time.Since(v.CreateTime()))-m.activeTime>0{
						delete(m.sessions,k)
					}
				}
		}

	}
}

func sessionId()string{
	return uuid.NewV4().String()
}


func (s *session)put(k,v interface{}){
	s.mux.Lock()
	defer s.mux.Unlock()
	s.element[k]=v
}
func (s *session)get(k interface{})interface{}{
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.element[k]
}
func (s *session)reset(){
	s.timeMux.Lock()
	defer s.timeMux.Unlock()
	s.createdTime=time.Now()
}
func (s *session)CreateTime()time.Time{
	s.timeMux.RLock()
	defer s.timeMux.RUnlock()
	return s.createdTime
}