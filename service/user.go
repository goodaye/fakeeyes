package service

import (
	"fmt"
	"time"

	"github.com/goodaye/fakeeyes/dao/rdb"

	"github.com/gorilla/websocket"
	"github.com/zheng-ji/goSnowFlake"

	"github.com/goodaye/fakeeyes/protos/request"
	uuid "github.com/satori/go.uuid"
)

var UIDWorker *goSnowFlake.IdWorker

func init() {
	var err error
	UIDWorker, err = goSnowFlake.NewIdWorker(1)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Entity
	rdb.User
	rdb.UserSession
}

// 登陆
func Login(req request.UserLogin) (user User, err error) {

	var dbuser rdb.User

	var has bool
	session := rdb.NewSession()
	defer session.Close()

	has, err = session.Where("name = ?", req.Name).Get(&dbuser)
	if err != nil {
		return
	}
	if !has {
		err = ErrorUserNotFound
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

// 通过token 验证用户
func LoginByToken(token string) (user *User, err error) {

	var dbuser rdb.User
	var dbusersession rdb.UserSession
	var has bool
	user = &User{}
	session := rdb.NewSession()
	defer session.Close()

	has, err = session.Where("token = ?", token).Get(&dbusersession)
	if err != nil {
		return
	}
	if !has {
		err = ErrorUserNotFound
		return
	}
	has, err = session.ID(dbusersession.UserID).Get(&dbuser)
	if err != nil {
		return
	}
	if !has {
		err = ErrorUserNotFound
		return
	}
	user.User = dbuser
	user.Session = session
	user.UserSession = dbusersession
	return
}

// 注册
func UserSignUp(req request.UserSignUp) (user User, err error) {

	var dbuser rdb.User

	var has bool
	var uid int64
	session := rdb.NewSession()
	defer session.Close()

	has, err = session.Where("name = ?", req.Name).Get(&dbuser)
	if err != nil {
		return
	}
	if has {
		err = ErrorUserExist
		return
	}
	for i := 0; i < 3; i++ {

		uid, err = GetUID()
		if err != nil {
			continue
		}
		newuser := rdb.User{
			UID:       fmt.Sprint(uid),
			Name:      req.Name,
			LastLogin: time.Now(),
		}
		_, err = session.Insert(&newuser)
		if err != nil {
			return
		}
		break
	}
	if err != nil {
		return
	}
	session.Commit()
	return
}

// 创建token
func (user *User) CreateToken() (err error) {

	var dbsession rdb.UserSession
	session := user.WithSession()
	defer session.Close()

	has, err := session.Where("user_id = ? ", user.User.ID).Get(&dbsession)
	if err != nil {
		return err
	}
	token := GenToken()
	if !has {
		// 创建新的dbsssion
		newdbss := rdb.UserSession{
			UserID:     user.User.ID,
			Token:      token,
			ExpireTime: time.Now().Add(UserTokenExpireDuration),
		}
		_, err = session.Insert(newdbss)
		if err != nil {
			session.Rollback()
			return
		}
	} else {
		// 更新现有的
		updatedbss := rdb.UserSession{
			Token:      token,
			ExpireTime: time.Now().Add(UserTokenExpireDuration),
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
	session.Commit()
	user.UserSession = dbsession
	return
}

func (u *User) ConnectDevice(req request.ConnectDevice, conn *websocket.Conn) (nroom *Room, err error) {

	nroom, err = CreateRoom(u, conn, req.DeviceUUID)
	if err != nil {
		return
	}
	nroom.Run()
	return

}

func GenToken() string {
	u4 := uuid.NewV4()
	return u4.String()

}

// ListDevices 列举设备列表
func (user *User) ListDevices() ([]rdb.Device, error) {

	var dbdevice []rdb.Device
	var err error
	session := user.WithSession()
	defer session.Close()
	err = session.Find(&dbdevice)
	return dbdevice, err

}

func GetUID() (uid int64, err error) {

	for i := 0; i < 3; i++ {
		uid, err = UIDWorker.NextId()
		if err != nil {
			return
		}
	}
	return
}
