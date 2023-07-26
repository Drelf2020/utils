package request

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var ErrJsonParse = errors.New("解析 JSON 失败")

type M map[string]string

func (m *M) Scan(val any) error {
	if val, ok := val.([]byte); ok {
		return json.Unmarshal(val, m)
	}
	return ErrJsonParse
}

func (m M) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// 请求任务
type Job struct {
	// GET 或 POST
	Method string `form:"method" yaml:"method" json:"method"`
	// 请求网址
	Url string `form:"url" yaml:"url" json:"url"`
	// 附带数据
	Data M `form:"data" yaml:"data" json:"data"`
	// 请求头
	Headers M `form:"headers" yaml:"headers" json:"headers"`
}

// 发送请求
func (job *Job) Request() *Result {

	// 添加 POST 参数
	ploady := make(url.Values)
	if job.Method == "POST" {
		AddAll(ploady, job.Data)
	}

	// 新建请求
	req, _ := http.NewRequest(job.Method, job.Url, strings.NewReader(ploady.Encode()))

	// 添加 GET 参数
	if job.Method == "GET" {
		q := req.URL.Query()
		AddAll(q, job.Data)
		req.URL.RawQuery = q.Encode()
	}

	// 添加请求头
	AddAll(req.Header, job.Headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 新建客户端
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return &Result{resp.StatusCode, body}
}

// 构造函数
func NewJob(method, url string, options ...Option) *Job {
	job := Job{
		method,
		url,
		make(map[string]string),
		make(map[string]string),
	}
	for _, op := range options {
		op(&job)
	}
	return &job
}
