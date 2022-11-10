package sqldev

import (
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
		Name:     "账号必填111",
		UserName: "姓名必填111",
		Pwd:      "mimabitian111@#¥",
		Expire:   -1,
		Role:     9,
		Sex:      1,
		Email:    "bitian@youxiang.com",
		Telno:    "13966661234",
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", a)
}

func TestGetUserInfo(t *testing.T) {
	a, err := sqldev.GetUserInfo(0)
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
