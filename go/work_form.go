package sqldev

type AuditWorkForm struct {
	Id           string `json:"id"`            //名称
	Status       int    `json:"status"`        //类型 1 通过  2 不通过 3 通过并执行
	CancelRemark string `json:"cancel_remark"` //不通过的理由
}
