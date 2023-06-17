package utils

// 这垃圾语言怎么连 bool 异或都没有啊
func Xor(x, y bool) bool {
	return (x && !y) || (!x && y)
}

// 全对
func All(expr ...bool) bool {
	for _, e := range expr {
		if !e {
			return false
		}
	}
	return true
}

// 过滤数组
func Filter[T any](v []T, f func(T) bool) (r []T) {
	for _, o := range v {
		if f(o) {
			r = append(r, o)
		}
	}
	return
}

// 类似 python 的 map 函数
func Map[T any, V any](f func(T) V, iter []T) (v []V) {
	for _, i := range iter {
		v = append(v, f(i))
	}
	return
}

// 条件遍历数组
//
// 好牛逼的函数
func ForEach[T any](v []T, f func(T), options ...func(T) bool) {
	Map(func(o T) int {
		if All(Map(func(t func(T) bool) bool { return t(o) }, options)...) {
			f(o)
		}
		return 0
	}, v)
}

// 复制字典
func CopyMap[T any](originalMap map[string]T) map[string]T {
	// Create the target map
	targetMap := make(map[string]T)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}
	return targetMap
}

// 过滤字典
//
// 注意 这会改变字典内容 使用前请复制一份
func FilterMap[T any](v map[string]T, f func(string, T) bool) {
	for k, o := range v {
		if !f(k, o) {
			delete(v, k)
		}
	}
}

// 条件遍历字典
//
// 该方法内部会复制一份字典
func ForMap[T any](v map[string]T, f func(string, T), options ...func(string, T) bool) {
	v = CopyMap(v)
	for _, option := range options {
		FilterMap(v, option)
	}
	for k, o := range v {
		f(k, o)
	}
}
