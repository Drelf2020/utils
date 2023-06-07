package request

// 简化 Get 请求
func Get(url string, options ...Option) *Result {
	return NewJob("GET", url, options...).Request()
}

// 简化 Post 请求
func Post(url string, options ...Option) *Result {
	return NewJob("POST", url, options...).Request()
}
