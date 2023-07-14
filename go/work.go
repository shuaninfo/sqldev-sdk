package sqldev

import (
	"encoding/json"
	"errors"
	"sqldev/utils"
	"time"
)

type WorkLog struct {
	ID      int64  `gorm:"primary_key" json:"id" db:"id"`
	WorkId  string `gorm:"type:varchar(128);index:idx_key_work_log_wid;" json:"work_id" form:"work_id"` //数据库id
	Type    int    `json:"type" form:"type"`                                                            // 工单类型
	OptId   int64  `json:"opt_id" db:"opt_id" gorm:"type:bigint(9);"  `                                 //操作人
	OptMark string `gorm:"type:varchar(500);"  json:"opt_mark" form:"opt_mark"`
	OptIp   string `gorm:"type:varchar(64);"  json:"opt_ip" form:"opt_ip"`
	OptDesc string `gorm:"type:varchar(32);"  json:"opt_desc" form:"opt_desc"`
	OptInfo string `gorm:"type:varchar(1000);"  json:"opt_info" form:"opt_info"`

	CreatedAt  time.Time `gorm:"type:datetime(3)"  json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime(3)"  json:"updated_at" db:"updated_at"`
	CreateName string    `json:"create_name"`

	Ok    string `json:"ok"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type WorkFlowTpl struct {
	Desc    string   `json:"desc"`
	Auditor []string `json:"auditor"`
	Type    string   `json:"type"`
}

type TplTypes []*WorkFlowTpl

// WorkDto 工单信息Dto
type WorkDto struct {
	ID string `json:"id"`

	Name string `json:"name"` //工单名称

	//发生错误后选择继续还是终止，是否回滚已提交的数据	1 终止执行不回滚错误 2 继续执行并忽略错误  3 终止并回滚错误
	ErrRollback int `json:"err_rollback"`

	HasDel int `json:"has_del"`

	DbId   string `json:"db_id"` //数据库id
	DbName string `json:"db_name"`

	Project     int64  `json:"project"`      //项目id
	ProjectName string `json:"project_name"` //工单名称
	// 数据库类型
	DbType string `json:"db_type"`

	Owner string `json:"owner"` //库名

	Schema string `json:"schema"` //库名

	Tab string `json:"tab"` //表名

	Status int `json:"status"` // '操作类型,0提交/待审核、1审核通过、2审核不通过、3 人工终止流程、4 定时执行、5执行中、6执行成功、7 执行失败、8 执行成功（线下）、9 系统异常终止',

	ExecStatus int `json:"exec_status"` //执行状态 0 待执行  1 执行中 2 定时执行、3 执行中、4 执行成功 5 执行失败 6 执行成功（线下） 7 系统异常终止

	AuditToken string `json:"audit_token"`

	HasToken bool `json:"has_token"` //支持线下审核

	Type int `json:"type"` //类型   1 修改(DML)  2 修改(DDL)  3 查询(DQL)  4 查询(导出)  5 其他脚本

	SqlType int `json:"sql_type"` // 工单详细的类型  12 update 13 insert 14 delete 15 ddl

	CreateId int64 `json:"create_id"` //发起人

	CreateIp string `json:"create_ip"` //发起人ip

	AuditId int64 `json:"audit_id"` //最终执行人

	FetchSize int64 `json:"fetch_size"` //查询大小

	//短信校验失败次数
	ErrCount int64 `json:"err_count"`

	FileType int `json:"file_type"` //查询大小

	WorkManual bool `json:"work_manual"` //支持线下审核

	NoMaskField string `json:"no_mask_field"` //不脱敏的字段

	BackUp int ` json:"back_up"` // 是否备份  0 不备份  1 备份

	ExecFlag int `json:"exec_flag"` // 执行时间  1 不限  2 指定范围

	ExecStart time.Time `json:"exec_start"` //执行开始时间

	ExecEnd time.Time `json:"exec_end"` //执行结束时间

	ConsumeTime int64 `json:"consume_time"` //消耗时间

	SqlAuditAuth string `json:"sql_audit_auth"` //Sql审计流程,默认当前组审核

	CurAudit int64 ` json:"cur_audit"` //是否有当前的审核组

	NextAudit int64 ` json:"next_audit"` //是否有下一个审核组

	FlowIndex int64 `json:"flow_index"` //当前的审核组下标,更新det

	//这里启用新的审核流程
	Steps TplTypes `json:"steps"`

	CurAuditUser string `json:"cur_audit_user"` //当前的审核用户

	CurAuditIdStr string `json:"cur_audit_id_str"` //当前的审核组id

	CurAuditType string `json:"cur_audit_type"` // 1 申请 2 审核 3 执行

	StartTime time.Time `json:"start_time"` //工单开始时间
	EndTime   time.Time `json:"end_time"`   //工单结束时间

	Result string `json:"result"`

	ExecResult string `json:"exec_result"`

	Mark string `json:"mark"`

	//支持短信验证码审核
	SmsAudit bool `json:"sms_audit"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 操作类型，1只读 2.读写 4.管理

	Sql string `json:"sql"`

	NextAuditIdx int `json:"next_audit_idx"` //当前的审核组下标

	ProjectNameMap map[int64]string `json:"project_name_map"` //工单名称

	SqlAuditAuthStr string ` json:"sql_audit_auth_str"` //Sql审计流程,默认当前组审核

	SqlAuditAuthArr []int64 ` json:"sql_audit_auth_arr"` //Sql审计流程,默认当前组审核

	AffectedRows int64 `json:"affected_rows"` //受影响行数(预估)affected_rows

	SqlHash string `json:"sql_hash"` //sql的哈希值,sqlsha1,用在pt等DDL工具上

	HasAuth bool `json:"has_auth"`

	CanView   bool `json:"can_view"`   //当前用户是否有查看工单权限
	CanReview bool `json:"can_review"` //当前用户是否有权限审核
	CanCancel bool `json:"can_cancel"` //当前用户是否有权限终止，只有审核人和提交人可终止，审核通过但未执行的工单，有执行权限的用户终止
	CanExec   bool `json:"can_exec"`   //当前用户是否有权限执行，只有审核人和提交人可终止，审核通过但未执行的工单，有执行权限的用户终止

	//history查询是不使用模糊匹配
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
	OrderBy    int    `json:"order_by"`
	CreateName string `json:"create_name"`

	FlowName   string `json:"flow_name"`   //当前的审核组下标,更新det
	FlowType   string `json:"flow_type"`   //当前的审核组下标,更新det
	FlowAction string `json:"flow_action"` //当前的审核组下标,更新det
	FlowStatus int    `json:"flow_status"` //

	NextAuditIdStr string `json:"next_audit_str_id"`  //当前的审核组id
	NextAuditUser  string ` json:"next_audit_user"  ` //下一个审核用户
	NextAuditType  string `json:"next_audit_type"`    // 1 申请 2 审核 3 执行

	//审核状态
	AuditStatus int      `json:"audit_status"`
	ProjectId2  int64    `json:"project_id2"`
	Ok          string   `json:"ok"`
	Start       string   `json:"start"`
	End         string   `json:"end"`
	WorkUser    *UserDto `json:"work_user"`

	LogList []*WorkLog `json:"log_list"`
}

// GetWorkInfo 获取工单信息
func (s *Sqldev) GetWorkInfo(id string) (*WorkDto, error) {
	params := map[string]string{
		"id": id,
	}

	res, err := s.sendRequest("POST", "/work/info", params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int      `json:"result"`
		Msg    string   `json:"msg"`
		Data   *WorkDto `json:"data"`
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

// AuditWork 工单审批
func (s *Sqldev) AuditWork(form *AuditWorkForm) error {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return err
	}

	res, err := s.sendRequest("POST", "/work/audit", params)
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
