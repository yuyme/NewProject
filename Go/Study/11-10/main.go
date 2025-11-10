package main

import (
	"fmt"
	"sync"
	"time"
)

type UserSession struct {
	UserID         string
	LoginTime      time.Time
	LastActiveTime time.Time
	IPAddress      string
}

type SessionManger struct {
	sessions map[string]*UserSession
	mutex    sync.RWMutex
}

func newSessionManger() *SessionManger {
	return &SessionManger{
		sessions: make(map[string]*UserSession),
	}
}

func (sm *SessionManger) AddSession(token string, userid string, ipaddress string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	session := &UserSession{
		UserID:         userid,
		LoginTime:      time.Now(),
		LastActiveTime: time.Now(),
		IPAddress:      ipaddress,
	}

	sm.sessions[token] = session

}

func (sm *SessionManger) getSession(token string) (*UserSession, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	session, ok := sm.sessions[token]

	if ok {
		session.LastActiveTime = time.Now()
	}
	return session, ok
}

func (sm *SessionManger) RemoveSession(token string) {
	sm.mutex.Lock()
	sm.mutex.Unlock()

	delete(sm.sessions, token)
}

func (sm *SessionManger) CleanupExpiredSession(timeout time.Duration) {
	sm.mutex.Lock()
	sm.mutex.Unlock()

	now := time.Now()
	for token, session := range sm.sessions {
		if now.Sub(session.LastActiveTime) > timeout {
			delete(sm.sessions, token)
		}
	}
}

func main() {
	//创建一个用户会话管理器对象
	manger := newSessionManger()
	manger.AddSession("token1", "001", "192.168.56.49")

	if session, ok := manger.getSession("token1"); ok {
		fmt.Printf("查询到用户会话信息：%s, 登陆时间为：%v \n", session.UserID, session.LoginTime.Format(time.DateTime))
	} else {
		fmt.Println("未查询到用户会话信息")
	}
	//清理超时会话
	manger.CleanupExpiredSession(time.Hour)
	//移除会话
	manger.RemoveSession("token1")

}
