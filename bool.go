package utils

// 这垃圾语言怎么连 bool 异或都没有啊
func Xor(x, y bool) bool {
	return (x && !y) || (!x && y)
}

// 与门
func All(expr ...bool) bool {
	for _, e := range expr {
		if !e {
			return false
		}
	}
	return true
}

// 或门
func Any(expr ...bool) bool {
	for _, e := range expr {
		if e {
			return true
		}
	}
	return false
}
