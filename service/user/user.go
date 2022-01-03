package user

import (
	"time"

	"github.com/goodaye/fakeeyes/dao/rdb"
	"github.com/goodaye/fakeeyes/protos/db"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/service"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	service.Entity
	db.User
	db.UserSession
}

func Login(req request.UserLogin) (user User, err error) {

	var dbuser db.User

	var has bool
	session := rdb.NewSession()
	defer session.Close()

	has, err = session.Where("name = ?", req.Name).Get(&dbuser)
	if err != nil {
		return
	}
	if !has {
		err = service.ErrorUserNotFound
		return
	}
	user.User = dbuser
	user.Session = session
	err = user.CreateToken()
	if err != nil {
		return
	}
	return
}

func UserSignUp(req request.UserSignUp) (user User, err error) {

	var dbuser db.User

	var has bool
	session := rdb.NewSession()
	defer session.Close()

	has, err = session.Where("name = ?", req.Name).Get(&dbuser)
	if err != nil {
		return
	}
	if has {
		err = service.ErrorUserExist
		return
	}
	newuser := db.User{
		Name:      req.Name,
		LastLogin: time.Now(),
	}
	_, err = session.Insert(&newuser)
	if err != nil {
		return
	}
	session.Commit()
	return
}

func (user *User) CreateToken() (err error) {

	var dbsession db.UserSession
	session := user.WithSession()

	has, err := session.Where("user_id = ? ", user.User.ID).Get(&dbsession)
	if err != nil {
		return err
	}
	token := user.GenToken()
	if !has {
		// 创建新的dbsssion
		newdbss := db.UserSession{
			UserID:     user.User.ID,
			Token:      token,
			ExpireTime: time.Now().Add(service.UserTokenExpireDuration),
		}
		_, err = session.Insert(newdbss)
		if err != nil {
			session.Rollback()
			return
		}
	} else {
		// 更新现有的
		updatedbss := db.UserSession{
			Token:      token,
			ExpireTime: time.Now().Add(service.UserTokenExpireDuration),
		}
		_, err = session.ID(dbsession.ID).Update(&updatedbss)
		if err != nil {
			session.Rollback()
			return
		}
	}
	_, err = session.Where("user_id = ? ", user.User.ID).Get(&dbsession)
	if err != nil {
		return err
	}
	user.UserSession = dbsession
	return
}

func (user User) GenToken() string {
	u4 := uuid.NewV4()
	return u4.String()

}
