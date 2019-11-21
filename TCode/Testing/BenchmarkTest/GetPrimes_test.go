package BenchmarkTest

import "testing"

//使用以下语句做性能测试 -run=^$ 排除功能测试文件
//go test -bench=. -run=^$ ./BenchmarkTest

//4 个 CPU 执行次数 平均耗时
//BenchmarkGetPrimes-4      343524              3228 ns/op
func BenchmarkGetPrimes(b *testing.B) {
	//b.N 由1开始 1秒为上限 ，会自动增加
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}
