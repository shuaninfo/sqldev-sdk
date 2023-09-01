package sqldev

import (
	"encoding/json"
	"errors"
)

type SqlQueryDto struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
	Msg   string      `json:"msg"`
	Ok    bool        `json:"ok"`
	Sql   string      `json:"sql"`
	No    int         `json:"no"`
}

func (s *Sqldev) Query(form *SqlQueryForm) (*[]SqlQueryDto, error) {
	params, err := ObjectToMap(form)
	if err != nil {
		return nil, err
	}

	res, err := s.sendRequest("POST", "/sql/query", params)
	if err != nil {
		return nil, err
	}

	result := &struct {
		Result int            `json:"result"`
		Msg    string         `json:"msg"`
		Data   *[]SqlQueryDto `json:"data"`
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
