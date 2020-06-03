/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:36:11
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 10:50:27
 */

package dao_test

import (
	"fmt"
	"govideo_server/dao"
	"govideo_server/model"
	"testing"
)

var id uint

func TestMain(m *testing.M) {
	m.Run()
}

func TestSucceed(t *testing.T) {
	return
}

func TestUser(t *testing.T) {
	t.Run("测试清空用户表：", testDelUsers)
	t.Run("测试获取用户表：", testGetUsers)

	t.Run("测试添加用户：", testAddUser)
	t.Run("测试添加多个用户：", testAddUsers)
	t.Run("测试添加相同用户名：", testAddSameUsers)
	t.Run("测试获取用户表：", testGetUsers)

	t.Run("测试获取单用户：", testGetUser)
}

func testDelUsers(t *testing.T) {
	dao.DelUsers()
}

func testGetUsers(t *testing.T) {
	users := dao.GetUsers()
	for k, v := range users {
		fmt.Printf("users[%v]:%v\n", k, v)
	}
}

func testAddUser(t *testing.T) {
	user := &model.User{
		//ID:       0,
		Name:     "testAddUser",
		Password: "testAddUser111",
		Level:    0,
	}
	if !dao.AddUser(user) {
		t.Error("AddUser fail.")
	}
	id = user.ID
}

func testAddUsers(t *testing.T) {
	user1 := model.User{
		//ID:       0,
		Name:     "testAddUsers1",
		Password: "testAddUsers111",
		Level:    1,
	}
	user2 := model.User{
		//ID:       0,
		Name:     "testAddUsers2",
		Password: "testAddUsers222",
		Level:    2,
	}
	if !dao.AddUser(&user1) {
		t.Error("AddUser1 fail.")
	}
	if !dao.AddUser(&user2) {
		t.Error("AddUser2 fail.")
	}
}

func testAddSameUsers(t *testing.T) {
	user1 := model.User{
		//ID:       0,
		Name:     "testAddSameUsers",
		Password: "testAddSameUsers111",
		Level:    1,
	}
	user2 := model.User{
		//ID:       0,
		Name:     "testAddSameUsers",
		Password: "testAddSameUsers222",
		Level:    2,
	}
	if !dao.AddUser(&user1) {
		t.Error("AddUser1 fail.")
	}
	if dao.AddUser(&user2) {
		t.Error("AddUser2 fail.")
	}
}

func testGetUser(t *testing.T) {
	user := dao.GetUserByID(id)
	if user.ID == 0 {
		t.Error("GetUserByID fail.")
	}
	fmt.Printf("user: %v\n", user)

	user = dao.GetUserByName("testAddUser")
	if user.ID == 0 {
		t.Error("GetUserByName fail.")
	}
	fmt.Printf("user: %v\n", user)

	user = dao.GetUserByName("notexist")
	if user.ID != 0 {
		t.Error("GetUserByName for notexist fail.")
	}
}
