/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:36:11
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-08 11:16:17
 */

package dao_test

import (
	"fmt"
	"govideo_server/dao"
	"govideo_server/model"
	"testing"
)

var (
	uid int64
	vid int64
)

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

	t.Run("测试清空用户表：", testDelUsers)
}

func TestVideo(t *testing.T) {
	t.Run("测试清空视频表：", testDelVideos)
	t.Run("测试获取视频表：", testGetVideos)

	t.Run("测试添加视频：", testAddVideo)
	t.Run("测试添加多个视频：", testAddVideos)
	t.Run("测试获取视频表：", testGetVideos)

	t.Run("测试获取视频：", testGetVideo)

	t.Run("测试获取视频数：", testGetVideosCount)
	t.Run("测试清空视频表：", testDelVideos)
	t.Run("测试获取视频数：", testGetVideosCount)
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
	uid = user.ID
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
	user := dao.GetUserByID(uid)
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

func testDelVideos(t *testing.T) {
	dao.DelVideos()
}

func testGetVideos(t *testing.T) {
	videos := dao.GetVideos()
	for k, v := range videos {
		fmt.Printf("videos[%v]:%v\n", k, v)
	}
}

func testAddVideo(t *testing.T) {
	video := &model.Video{
		//ID:       0,
		Title:       "testAddVideo",
		Description: "testAddVideo111",
		Filepath:    "Filepath111",
		Userid:      1,
	}
	if !dao.AddVideo(video) {
		t.Error("AddVideo fail.")
	}
	vid = video.ID
}

func testAddVideos(t *testing.T) {
	video1 := model.Video{
		//ID:       0,
		Title:       "testAddVideos",
		Description: "testAddVideos111",
		Filepath:    "Filepath111",
		Userid:      1,
	}
	video2 := model.Video{
		//ID:       0,
		Title:       "testAddVideos",
		Description: "testAddVideos222",
		Filepath:    "Filepath222",
		Userid:      1,
	}
	if !dao.AddVideo(&video1) {
		t.Error("AddVideo1 fail.")
	}
	if !dao.AddVideo(&video2) {
		t.Error("AddVideo2 fail.")
	}
}

func testGetVideo(t *testing.T) {
	video := dao.GetVideoByID(vid)
	if video.ID == 0 {
		t.Error("GetVideoByID fail.")
	}
	fmt.Printf("video: %v\n", video)

	video = dao.GetVideoByTitle("testAddVideo")
	if video.ID == 0 {
		t.Error("GetVideoByTitle fail.")
	}
	fmt.Printf("video: %v\n", video)

	video = dao.GetVideoByTitle("notexist")
	if video.ID != 0 {
		t.Error("GetVideoByTitle for notexist fail.")
	}

	videos := dao.GetVideosByLimit(3, 1)
	if len(videos) == 0 {
		t.Error("GetVideosByLimit fail.")
	}
	for k, v := range videos {
		fmt.Printf("videos[%v]:%v\n", k, v)
	}
}

func testGetVideosCount(t *testing.T) {
	fmt.Println("videos count: ", dao.GetVideosCount())
}
