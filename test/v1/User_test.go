package v1

import (
	"materiel/src/api/v1/Users"
	"materiel/src/db/Schema"
	"materiel/src/util"
	"testing"
)

func Test_SearchUserByName(t *testing.T) {
	u := Users.FindUserByUsername("lvmingyin")
	if u.Nickname == "吕铭印" {
		t.Log("根据用户名查询成功")
	} else {
		t.Error("根据用户名查询失败")
	}
}

func Test_SearchUserById(t *testing.T) {
	u := Users.FindUserById(1)
	if u.Nickname == "吕铭印" {
		t.Log("根据id查询成功")
	} else {
		t.Error("根据id查询失败")
	}
}

func Test_UserAddQueryUpdateDelete(t *testing.T) {
	salt := util.GetRandomString()
	password := util.GetSha256Password("test", salt)
	u := Schema.User{Username: "test", Password: password, Nickname: "测试账号", Salt: salt}
	insertId, err := Users.AddUser(u)
	if err != nil {
		t.Error("新增用户错误", err)
		return
	}
	t.Log("新增用户成功:", insertId)
	u.User_id = insertId

	// 修改密码
	u.Password = util.GetSha256Password("testnew", salt)

	affectedRow, err := Users.UpdateUser(u)
	if err != nil {
		t.Error("更新用户错误", err)
		return
	}
	t.Log("更新用户成功:", affectedRow)

	u = Users.FindUserById(u.User_id)
	if u.Username == "test" {
		t.Log("查询用户成功", u)
	} else {
		t.Error("查询用户错误", err)
	}

	affectedRow, err = Users.DeleteUserById(u.User_id)
	if err != nil {
		t.Error("删除失败")
	} else {
		t.Log("删除成功", affectedRow)
	}
}
