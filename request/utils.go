package request

// 将 data 数据一股脑加进 vm
func AddAll(vm interface{ Add(key, value string) }, data map[string]string) {
	for k, v := range data {
		vm.Add(k, v)
	}
}
