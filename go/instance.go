package sqldev

import (
	"encoding/json"
	"errors"
	"sqldev/utils"
	"time"
)

// InstanceDto 数据源信息Dto
type InstanceDto struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`

	User string `json:"user"`

	Pass string `json:"pass"`

	Ip string `json:"ip"`

	Port string `json:"port"`

	Db string `json:"db"`

	SshUse      bool   ` json:"ssh_use"`
	SshServer   string `json:"ssh_server"`
	SshPort     string `json:"ssh_port"`
	SshUser     string `json:"ssh_user"`
	SshPassword string `json:"ssh_password"`
	SshAuth     int    `json:"ssh_auth"` //1 password 2 ssh privateKey

	//连接类型 TNS 或 空值
	ConType string `json:"con_type"`

	// Orcale服务类型 service或sid
	ServiceType string `json:"service_type"`

	// 普通用户还是dba
	UserType string `json:"user_type"`

	// 数据库类型
	DbType string `json:"db_type"`
	// 操作类型，1只读 2.读写 4.管理
	OperType int `json:"oper_type"`

	Tty bool `json:"tty" db:"tty"` // web控制台 0 关闭 1 开启

	//标签
	Tags string `json:"tags"`

	//是否支持事务
	OpenTx bool ` json:"open_tx"`

	PoolMin int `json:"pool_min"`

	PoolMax int `json:"pool_max"`

	//最大获取行数
	MaxRow int64 `json:"max_row"`

	//手动事务超时时间，单位秒，默认30秒
	TxTimeout int64 `json:"tx_timeout"`

	//是否支持异步查询
	AsynQuery int `json:"asyn_query"`

	//是否支持文件导出
	AsynExport int `json:"asyn_export"`

	// 是否支持sql转储
	AsynSql int `json:"asyn_sql"`

	AddUser int64 ` json:"add_user"`

	//分页查询方式 1 包装limit    2 游标
	PageQuery int `json:"page_query"`

	//使用状态 1 正常 0停用
	State int `json:"state"`

	ConState int `json:"con_state"` //连接状态，3 连接失败 2 已连接  1 连接中 0 未连接

	//拓展配置
	Config string `json:"config"`

	//连接错误的错误信息
	ErrMsg string `json:"err_msg"`
}

// GetInstanceInfo 获取数据源信息
func (s *Sqldev) GetInstanceInfo(id string) (*InstanceDto, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/instance/info", params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int          `json:"result"`
		Msg    string       `json:"msg"`
		Data   *InstanceDto `json:"data"`
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

// GetInstancePage 获取数据源信息列表
func (s *Sqldev) GetInstancePage(form *InstancePageForm) ([]*InstanceDto, int64, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return nil, 0, err
	}

	res, err := s.sendRequest("POST", "/instance/page", params)
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

// AddInstance 添加数据源
func (s *Sqldev) AddInstance(form *InstanceAddForm) (string, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return "", err
	}

	res, err := s.sendRequest("POST", "/instance/add", params)
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

// UpdateInstance 更新数据源信息
func (s *Sqldev) UpdateInstance(form *InstanceUpdForm) (string, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return "", err
	}

	res, err := s.sendRequest("POST", "/instance/upd", params)
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

// RemoveInstance 删除数据源
func (s *Sqldev) RemoveInstance(id string) error {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/instance/remove", params)
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

// StateInstance 启用/禁用数据源
func (s *Sqldev) StateInstance(form *InstanceStateForm) error {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return err
	}

	res, err := s.sendRequest("POST", "/instance/state", params)
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

// GetInstanceSync 同步数据源
func (s *Sqldev) GetInstanceSync() ([]*InstanceDto, error) {

	res, err := s.sendRequest("GET", "/instance/sync", nil)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int            `json:"result"`
		Msg    string         `json:"msg"`
		Data   []*InstanceDto `json:"data"`
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
