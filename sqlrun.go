package sqldev

import (
	"encoding/json"
	"errors"
	"sqldev/utils"
)

type SqlQueryForm struct {
	Project   int64  `json:"project"`
	Iid       string `json:"iid"`
	Sql       string `json:"sql"`
	Owner     string `json:"owner"`
	Schema    string `json:"schema"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
	NeedTotal bool   `json:"need_total"`
}

type SqlQueryDto struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func (s *Sqldev) Query(form *SqlQueryForm) (interface{}, error) {
	params, err := utils.ObjectToMap(form)
	if err != nil {
		return nil, err
	}

	res, err := s.sendRequest("POST", "/sql/query", params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int          `json:"result"`
		Msg    string       `json:"msg"`
		Data   *SqlQueryDto `json:"data"`
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
