package sqldev

// InstancePageForm 实例列表查询参数
type InstancePageForm struct {
	PageNo    int    `json:"page_no"`
	PageSize  int    `json:"page_size"`
	Name      string `json:"name"`
	Ip        string `json:"ip"`
	GroupId   string `json:"group_id"`
	State     string `json:"state"`
	DbType    string `json:"db_type"`
	ConState  string `json:"con_state"`
	CronState string `json:"cron_state"`
	OrderBy   int    `json:"order_by"`
	TagId     int64  `json:"tag_id"`
}

// InstanceAddForm 添加实例参数
type InstanceAddForm struct {
	ID string `json:"id"`

	Name string `json:"name"`

	User string `json:"user"`

	Pass string ` json:"pass"`

	Ip string ` json:"ip"`

	Port string ` json:"port"`

	Db string ` json:"db"`

	ConType string `json:"con_type"`

	ServiceType string `json:"service_type"`

	UserType string `json:"user_type"`

	OperType int `json:"oper_type"`

	DbType string `json:"db_type"`

	AsynExport int `json:"asyn_export"`

	AsynQuery int `json:"asyn_query"`

	AsynSql int `json:"asyn_sql"`

	PageQuery int `json:"page_query"`

	Config string ` json:"config"`

	State int ` json:"state"`

	PoolMin int `json:"pool_min"`

	PoolMax int `json:"pool_max"`

	MaxRow int `json:"max_row"`

	SshUse      bool   ` json:"ssh_use"`
	SshServer   string ` json:"ssh_server"`
	SshPort     string ` json:"ssh_port"`
	SshUser     string ` json:"ssh_user"`
	SshPassword string ` json:"ssh_password"`
	SshAuth     int    ` json:"ssh_auth"`

	//手动事务超时时间，单位秒，默认30秒
	TxTimeout int64 `json:"tx_timeout"`
}

// InstanceUpdForm 更新实例参数
type InstanceUpdForm struct {
	ID string `json:"id"`

	Name string `json:"name"`

	User string `json:"user"`

	Pass string ` json:"pass"`

	Ip string ` json:"ip"`

	Url string ` json:"url"`

	Port string ` json:"port"`

	Db string ` json:"db"`

	CharSet string ` json:"char_set"`

	DbVersion string ` json:"db_version"`

	ConType string `json:"con_type"`

	ServiceType string `json:"service_type"`

	UserType string `json:"user_type"`

	OperType int `json:"oper_type"`

	DbType string `json:"db_type"`

	AsynExport int `json:"asyn_export"`

	AsynQuery int `json:"asyn_query"`

	AsynSql int `json:"asyn_sql"`

	PageQuery int `json:"page_query"`

	Config string ` json:"config"`

	State int ` json:"state"`

	PoolMin int `json:"pool_min"`

	PoolMax int `json:"pool_max"`

	MaxRow int `json:"max_row"`

	Projects string `json:"projects"`

	TagIds string ` json:"tag_ids"`

	Tty bool `json:"tty"` // web控制台 0 关闭 1 开启 2 须走工单

	SshUse      bool   ` json:"ssh_use"`
	SshServer   string ` json:"ssh_server"`
	SshPort     string ` json:"ssh_port"`
	SshUser     string ` json:"ssh_user"`
	SshPassword string ` json:"ssh_password"`
	SshAuth     int    ` json:"ssh_auth"`

	//手动事务超时时间，单位秒，默认30秒
	TxTimeout int64 `json:"tx_timeout"`
}

// InstanceStateForm 启用/禁用实例参数
type InstanceStateForm struct {
	ID    string `json:"id"`
	State int    `json:"state"`
}
