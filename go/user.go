package sqldev

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// UserDto 用户信息Dto
type UserDto struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Role      int       `json:"role"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	State     int64     `json:"state"`
	Sex       int       `json:"sex"`
	Email     string    `json:"email"`
	LastIp    string    `json:"last_ip"`
	LastTime  time.Time `json:"last_time"`
	Telno     string    `json:"telno"`
	Expire    int64     `json:"expire"`
	GroupName string    `json:"group_name"`
}

// GetUserInfo 获取用户信息
func (s *Sqldev) GetUserInfo(id int64) (*UserDto, error) {
	params := map[string]string{
		`id`: strconv.FormatInt(id, 10),
	}

	res, err := s.sendRequest("POST", "/user/info", params)
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
	params, err := ObjectToMap(form)
	if err != nil {
		return nil, 0, err
	}

	res, err := s.sendRequest("POST", "/user/page", params)
	if err != nil {
		return nil, 0, err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   struct {
			Count    int64      `json:"count"`
			PageNo   int64      `json:"page_no"`
			PageSize int64      `json:"page_size"`
			List     []*UserDto `json:"list"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, 0, err
	}

	if result.Result != 1 {
		return nil, 0, errors.New(result.Msg)
	}

	return result.Data.List, result.Data.Count, nil
}

// AddUser 添加用户
func (s *Sqldev) AddUser(form *UserAddForm) (int64, error) {
	params, err := ObjectToMap(form)
	if err != nil {
		return 0, err
	}

	res, err := s.sendRequest("POST", "/user/add", params)
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

// UpdUser 更新用户信息
func (s *Sqldev) UpdUser(form *UserUpdForm) error {
	params, err := ObjectToMap(form)
	if err != nil {
		return err
	}

	res, err := s.sendRequest("POST", "/user/upd", params)
	if err != nil {
		return err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   string `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}

	if result.Result != 1 {
		return errors.New(result.Msg)
	}
	return nil
}

// RemoveUser 删除用户
func (s *Sqldev) RemoveUser(id int64) error {
	params := map[string]string{
		`id`: strconv.FormatInt(id, 10),
	}

	res, err := s.sendRequest("POST", "/user/remove", params)
	if err != nil {
		return err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   string `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}

	if result.Result != 1 {
		return errors.New(result.Msg)
	}
	return nil
}

// StateUser 启用/禁用用户
func (s *Sqldev) StateUser(form *UserStateForm) error {
	params, err := ObjectToMap(form)
	if err != nil {
		return err
	}

	res, err := s.sendRequest("POST", "/user/state", params)
	if err != nil {
		return err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   string `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}

	if result.Result != 1 {
		return errors.New(result.Msg)
	}
	return nil
}

// GetProjectUserList 获取用户项目列表
func (s *Sqldev) GetUserProjectList(name string) ([]*ProjectDto, int64, error) {
	params := map[string]string{
		"name": name,
	}

	res, err := s.sendRequest("POST", "/user/project/list", params)
	if err != nil {
		return nil, 0, err
	}

	result := &struct {
		Result int           `json:"result"`
		Msg    string        `json:"msg"`
		Data   []*ProjectDto `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, 0, err
	}

	if result.Result != 1 {
		return nil, 0, errors.New(result.Msg)
	}
	return result.Data, int64(len(result.Data)), nil
}
