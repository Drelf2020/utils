package request

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Drelf2020/utils"
)

// 请求任务
type Job struct {
	// GET 或 POST
	Method string
	// 请求网址
	Url string
	// 附带数据
	Data map[string]string
	// 请求头
	Headers map[string]string
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
	if utils.LogErr(err) {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if utils.LogErr(err) {
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
