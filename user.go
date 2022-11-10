package sqldev

import (
	"encoding/json"
	"errors"
	"sqldev/utils"
	"strconv"
	"time"
)

type UserDto struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Role       int       `json:"role"`
	Name       string    `json:"name"`
	UserName   string    `json:"user_name"`
	State      int64     `json:"state"`
	Sex        int       `json:"sex"`
	Email      string    `json:"email"`
	LastIp     string    `json:"last_ip"`
	LastTime   time.Time `json:"last_time"`
	Telno      string    `json:"telno"`
	Expire     int64     `json:"expire"`
	GroupId    string    `json:"group_id"`
	GroupName  string    `json:"group_name"`
	GroupPid   string    `json:"group_pid"`
	GroupPname string    `json:"group_pname"`
}

type UserAddForm struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
	Expire   int64  `json:"expire"`
	Role     int    `json:"role"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Telno    string `json:"telno"`
	RolesStr string `json:"roles_str"`
	GroupId  string `json:"group_id"`
}

type UserPageForm struct {
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	State    int64  `json:"state"`
	Role     int    `json:"role"`
	GroupId  string `json:"group_id"`

	RoleState string `json:"role_state"`
	Email     string `json:"email"`
	Telno     string `json:"telno"`
	Source    string `json:"source"`
	RoleId    int    `json:"role_id"`
	OrderBy   int    `json:"order_by"`
}

// AddUser 添加用户
func (s *Sqldev) AddUser(form *UserAddForm) (int64, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return 0, err
	}

	res, err := s.sendRequest(`POST`, `/user/add`, params)
	if err != nil {
		return 0, err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   int64  `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return 0, err
	}

	if result.Result != 1 {
		return 0, errors.New(result.Msg)
	}
	return result.Data, nil
}

// GetUserInfo 获取用户信息
func (s *Sqldev) GetUserInfo(id int64) (*UserDto, error) {
	params := map[string]string{
		`id`: strconv.FormatInt(id, 10),
	}

	res, err := s.sendRequest(`POST`, `/user/info`, params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int      `json:"result"`
		Msg    string   `json:"msg"`
		Data   *UserDto `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}

	if result.Result != 1 {
		return nil, errors.New(result.Msg)
	}

	return result.Data, nil
}

// GetUserPage 获取用户列表
func (s *Sqldev) GetUserPage(form *UserPageForm) ([]*UserDto, int64, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return nil, 0, err
	}

	res, err := s.sendRequest(`POST`, `/user/page`, params)
	if err != nil {
		return nil, 0, err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   struct {
			Data  []*UserDto `json:"data"`
			Total int64      `json:"total"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, 0, err
	}

	if result.Result != 1 {
		return nil, 0, errors.New(result.Msg)
	}

	return result.Data.Data, result.Data.Total, nil
}
