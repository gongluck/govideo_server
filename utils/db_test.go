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

func TestAddUser(t *testing.T) {
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

func TestAddUsers(t *testing.T) {
	user := &model.User{
		//ID:       0,
		Name:     "gongluck",
		Password: "123456",
		Level:    3,
	}
	if !AddUser(user) {
		t.Error("AddUser fail.")
	}
	if !AddUser(user) {
		t.Error("AddUser fail.")
	}
	if !AddUser(user) {
		t.Error("AddUser fail.")
	}
}
