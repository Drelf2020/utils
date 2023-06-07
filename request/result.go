package request

import "encoding/json"

// 结果
type Result struct {
	// 状态码
	Code int
	// 内容
	Data []byte
}

// 解析结果为文本
func (res Result) Text() string {
	return string(res.Data)
}

// 解析结果为 json
func (res Result) Json(data any) error {
	return json.Unmarshal(res.Data, data)
}
