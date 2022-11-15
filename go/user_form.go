package sqldev

// UserPageForm 用户列表查询参数
type UserPageForm struct {
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	State    int64  `json:"state"`
	Role     int    `json:"role"`
	GroupId  string `json:"group_id"`

	Email   string `json:"email"`
	Telno   string `json:"telno"`
	Source  string `json:"source"`
	RoleId  int    `json:"role_id"`
	OrderBy int    `json:"order_by"`
}

// UserAddForm 添加用户参数
type UserAddForm struct {
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Pwd       string `json:"pwd"`
	Expire    int64  `json:"expire"`
	Role      int    `json:"role"`
	Sex       int    `json:"sex"`
	Email     string `json:"email"`
	Telno     string `json:"telno"`
	RolesStr  string `json:"roles_str"`
	GroupName string `json:"group_name"`
}

// UserUpdForm 更新用户信息参数
type UserUpdForm struct {
	Id        int64  `json:"id"`
	GroupName string `json:"group_name"`

	State    int64  `json:"state"`
	Expire   int64  `json:"expire"`
	TelNo    string `json:"tel_no"`
	RolesStr string `json:"roles_str"`
	UserName string `json:"user_name"`
	Role     int    `json:"role"`
	Telno    string `json:"telno"`
	Email    string `json:"email"`
	Sex      int    `json:"sex"`
}

// UserStateForm 启用/禁用用户参数
type UserStateForm struct {
	Id    int64 `json:"id"`
	State int64 `json:"state"`
}
