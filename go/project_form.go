package sqldev

// ProjectPageForm 项目列表查询参数
type ProjectPageForm struct {
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
}

// ProjectAddForm 添加项目参数
type ProjectAddForm struct {
	Name string `json:"name" `

	Icon string `json:"icon"`

	Expire int64 `json:"expire"`

	PublicState int ` json:"public_state"` // 1 否  2 是
}

// ProjectUpdForm 更新项目参数
type ProjectUpdForm struct {
	ID string `json:"id"`

	Name string `json:"name" `

	Icon string `json:"icon"`

	Expire int64 `json:"expire"`

	PublicState int ` json:"public_state"` // 1 否  2 是
}

// ProjectStateForm 启用/禁用项目参数
type ProjectStateForm struct {
	ID    string `json:"id"`
	State int    `json:"state"`
}
