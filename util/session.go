/*
 * @Author: gongluck
 * @Date: 2020-06-08 13:33:41
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 13:23:28
 */

package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	sessionid := uuid.NewV4()
	return sessionid.String()
}

func SetSession(c *gin.Context, userid int64) error {
	uuidstr := NewUUID()
	session := sessions.Default(c)
	session.Set("session", uuidstr)
	session.Set(uuidstr, userid)
	return session.Save()
}

func GetSessionAndUser(c *gin.Context) (interface{}, interface{}) {
	session := sessions.Default(c)
	vsession := session.Get("session")
	if vsession == nil {
		return nil, nil
	}

	user := session.Get(vsession)
	return vsession, user
}

func DelSession(c *gin.Context) error {
	session := sessions.Default(c)
	vsession, user := GetSessionAndUser(c)

	if vsession != nil {
		session.Delete(vsession)
	}
	if user != nil {
		session.Delete(user)
	}
	return session.Save()
}

func GetUserID(c *gin.Context) int64 {
	session := sessions.Default(c)
	vsession := session.Get("session")
	user := session.Get(vsession)
	if user == nil {
		return 0
	} else {
		return user.(int64)
	}
}
