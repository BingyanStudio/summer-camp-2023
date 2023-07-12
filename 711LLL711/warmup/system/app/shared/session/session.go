package mysession

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

const SESSION_NAME = "session-name"

type Session struct {
	Id   string
	Role string
}

// 从session中获取id,role
func GetSession(req *http.Request) (*Session, error) {
	session, err := store.Get(req, SESSION_NAME)
	if err != nil {
		return nil, err
	}
	//如果session是新的，原来没有，返回空
	if session.IsNew {
		return nil, nil
	}

	id, ok := session.Values["id"].(string)

	//字段获取失败，返回空
	if !ok {
		return nil, nil
	}
	role, ok := session.Values["role"].(string)
	if !ok {
		return nil, nil
	}

	return &Session{
		Id:   id,
		Role: role,
	}, nil
}

func SetSession(userId, userRole string, req *http.Request, resp http.ResponseWriter) error {
	session, _ := store.Get(req, SESSION_NAME)
	session.Values["id"] = userId
	session.Values["role"] = userRole
	log.Println("session saved: ", session.Values["id"], session.Values["role"])
	return session.Save(req, resp)
}
