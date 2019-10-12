package session

import (
	"fmt"
	"testing"
)

func TestSessionManager_ReadSession(t *testing.T) {
	sess:=DefalutSessionManger.CreateSession()
	b:=DefalutSessionManger.ReadSession(sess.sessionId)
	fmt.Println(b==sess)
}
