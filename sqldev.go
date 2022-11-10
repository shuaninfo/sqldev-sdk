package sqldev

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sqldev/sign"
	"strconv"
	"strings"
	"time"
)

// HTTPClient 可以直接使用标准库 net/http 中的 http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Sqldev 是一个用于调用 Sqldev API 的客户端结构体
type Sqldev struct {
	Endpoint  string // 接口地址
	AppKey    string //应用标识
	AppSecret string //应用密钥
	Client    HTTPClient
}

// NewSqldev 创建一个新的 Sqldev 客户端，需要传入应用标识和应用密钥
// 应用标识和应用密钥可以在接口文档中找到获取方式
// 如果不传入 HTTPClient，则默认使用标准库中的 http.Client
func NewSqldev(endpoint, appKey, appSecret string, client HTTPClient) *Sqldev {
	if client == nil {
		client = http.DefaultClient
	}
	sqldev := &Sqldev{
		endpoint,
		appKey,
		appSecret,
		client,
	}
	return sqldev
}

// sendRequest 负责发送请求
// method 是请求方法，可以是 GET 或 POST
// path 是请求路径，例如 /openapi/v1/user/info
// payload 是请求参数，会自动添加 app_key 和 current_time 参数
// 返回值是响应体的字节数组，如果出错则返回错误信息
func (s *Sqldev) sendRequest(method, path string, payload map[string]string) ([]byte, error) {
	params := url.Values{}
	for k, v := range payload {
		params.Set(k, v)
	}
	params.Set(`app_key`, s.AppKey)
	params.Set(`current_time`, strconv.FormatInt(time.Now().UnixMilli(), 10))

	sign, err := sign.GetSignature(params, s.AppSecret)
	if err != nil {
		return nil, err
	}

	params.Set(`sign`, sign)

	request, err := http.NewRequest(method, s.Endpoint+path, nil)
	if err != nil {
		return nil, err
	}

	if method == `GET` {
		request.URL.RawQuery = params.Encode()
	} else if method == `POST` {
		// POST 请求需要将参数放在请求体中，并设置 Content-Type 为 application/x-www-form-urlencoded
		request.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
		request.Body = ioutil.NopCloser(strings.NewReader(params.Encode()))
	} else {
		return nil, errors.New(`Unsupported method`)
	}

	response, err := s.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return nil, errors.New("Status code:" + strconv.Itoa(response.StatusCode) + "\nBody:" + string(body))
	}
	return body, nil
}
