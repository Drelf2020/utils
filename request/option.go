package request

type Option func(*Job)

// 数据
func Data(key, value string) Option {
	return func(job *Job) {
		job.Data[key] = value
	}
}

// 请求头
func Header(key, value string) Option {
	return func(job *Job) {
		job.Headers[key] = value
	}
}
