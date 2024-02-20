package sqldev

import (
	"encoding/json"
	"errors"
)

// ProjectDto 项目信息Dto
type ProjectDto struct {
	ID int64 `json:"id"`

	Name string `json:"name" `

	Icon string `json:"icon"`

	Expire int64 `json:"expire"`

	PublicState int ` json:"public_state"` // 1 否  2 是
}

// GetProjectInfo 获取项目信息
func (s *Sqldev) GetProjectInfo(id string) (*ProjectDto, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/project/info", params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int         `json:"result"`
		Msg    string      `json:"msg"`
		Data   *ProjectDto `json:"data"`
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

// GetProjectPage 获取项目信息列表
func (s *Sqldev) GetProjectPage(form *ProjectPageForm) ([]*ProjectDto, int64, error) {
	params, err := ObjectToMap(form)
	if err != nil {
		return nil, 0, err
	}

	res, err := s.sendRequest("POST", "/project/page", params)
	if err != nil {
		return nil, 0, err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   struct {
			Count    int64         `json:"count"`
			PageNo   int64         `json:"page_no"`
			PageSize int64         `json:"page_size"`
			List     []*ProjectDto `json:"list"`
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

// AddProject 添加项目
func (s *Sqldev) AddProject(form *ProjectAddForm) (string, error) {
	params, err := ObjectToMap(form)
	if err != nil {
		return "", err
	}

	res, err := s.sendRequest("POST", "/project/add", params)
	if err != nil {
		return "", err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   string `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return "", err
	}

	if result.Result != 1 {
		return "", errors.New(result.Msg)
	}
	return result.Data, nil
}

// UpdateProject 更新项目信息
func (s *Sqldev) UpdateProject(form *ProjectUpdForm) (string, error) {
	params, err := ObjectToMap(form)
	if err != nil {
		return "", err
	}

	res, err := s.sendRequest("POST", "/project/upd", params)
	if err != nil {
		return "", err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   string `json:"data"`
	}{}

	err = json.Unmarshal(res, result)
	if err != nil {
		return "", err
	}

	if result.Result != 1 {
		return "", errors.New(result.Msg)
	}
	return result.Data, nil
}

// RemoveProject 删除项目
func (s *Sqldev) RemoveProject(id string) error {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/project/remove", params)
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

// StateProject 启用/禁用项目
func (s *Sqldev) StateProject(form *ProjectStateForm) error {
	params, err := ObjectToMap(form)
	if err != nil {
		return err
	}

	res, err := s.sendRequest("POST", "/project/state", params)
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

// GetProjectUserList 获取项目成员列表
func (s *Sqldev) GetProjectUserList(id string) ([]*UserDto, int64, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/project/user/list", params)
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

// GetProjectUserList 获取项目数据源列表
func (s *Sqldev) GetProjectInstanceList(id string) ([]*InstanceDto, int64, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/project/instance/list", params)
	if err != nil {
		return nil, 0, err
	}

	result := &struct {
		Result int    `json:"result"`
		Msg    string `json:"msg"`
		Data   struct {
			Count    int64          `json:"count"`
			PageNo   int64          `json:"page_no"`
			PageSize int64          `json:"page_size"`
			List     []*InstanceDto `json:"list"`
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
