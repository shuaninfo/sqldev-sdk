package sqldev

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

const (
	ENDPOINT   = `http://127.0.0.1:8080/openapi/v1`
	APP_KEY    = `a9967fcfa669c40c5df7db31a33f70d4`
	APP_SECRET = `5414aabf-db54-4f45-b7e1-1d99aa10e47c`
)

var sqldev = NewSqldev(ENDPOINT, APP_KEY, APP_SECRET, &http.Client{})

func TestAddUser(t *testing.T) {
	a, err := sqldev.AddUser(&UserAddForm{
		Name:      "账号必填11122222233",
		UserName:  "姓名必填11222123121",
		Pwd:       "mimabitian111@#¥",
		Expire:    -1,
		Role:      9,
		Sex:       1,
		Email:     "bitian@youxiang.com",
		Telno:     "13966661234",
		GroupName: "自动分组11",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestUpdUser(t *testing.T) {
	err := sqldev.UpdUser(&UserUpdForm{
		Id:        17,
		Name:      "用户名222",
		GroupName: "自动创建！",
		State:     0,
		Expire:    -1,
		TelNo:     "",
		RolesStr:  "",
		UserName:  "名称222",
		Role:      9,
		Telno:     "",
		Email:     "",
		Sex:       0,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestRemoveUser(t *testing.T) {
	err := sqldev.RemoveUser(15)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestUpdateUserState(t *testing.T) {
	err := sqldev.StateUser(&UserStateForm{
		Id:    14,
		State: 1,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestGetUserInfo(t *testing.T) {
	a, err := sqldev.GetUserInfo(17)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestGetUserPage(t *testing.T) {
	a, total, err := sqldev.GetUserPage(&UserPageForm{
		PageNo:   0,
		PageSize: 100,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
	fmt.Printf("%+v\n", total)
}

func TestGetInstance(t *testing.T) {
	a, err := sqldev.GetInstance("3ab520c26bbd77f4b24db89aaaed71ec")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	b, _ := json.Marshal(a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

func TestGetInstancePage(t *testing.T) {
	a, total, err := sqldev.GetInstancePage(&InstancePageForm{
		PageNo:   0,
		PageSize: 100,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i := range a {
		fmt.Printf("%+v\n", a[i])
	}
	fmt.Printf("%+v\n", total)
}

func TestRemoveInstance(t *testing.T) {
	err := sqldev.RemoveInstance("5787240261629d40b4941ec4c85724e6")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
