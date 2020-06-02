/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:36:11
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-02 22:51:57
 */

package utils

import (
	"govideo_server/model"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestSucceed(t *testing.T) {
	return
}

func TestUser(t *testing.T) {
	t.Run("测试清空用户表：", testDelUsers)
	t.Run("测试添加用户：", testAddUser)
	t.Run("测试添加多个用户：", testAddUsers)
}

func testDelUsers(t *testing.T) {
	DelUsers()
}

func testAddUser(t *testing.T) {
	user := &model.User{
		//ID:       0,
		Name:     "gongluck",
		Password: "123456",
		Level:    0,
	}
	if !AddUser(user) {
		t.Error("AddUser fail.")
	}
}

func testAddUsers(t *testing.T) {
	user1 := model.User{
		//ID:       0,
		Name:     "gongluck",
		Password: "123456",
		Level:    1,
	}
	user2 := user1
	user2.Level = 2
	if !AddUser(&user1) {
		t.Error("AddUser1 fail.")
	}
	if !AddUser(&user2) {
		t.Error("AddUser2 fail.")
	}
}
