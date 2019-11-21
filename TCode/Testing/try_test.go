//此文件要以test结尾
package main

import (
	"math"
	"testing"
)

//单元测试 以Test开头
func TestAbs(t *testing.T) {
	var a, expect float64 = -110, 10
	// 只有当测试失败时才会显示 或者用 -v参数
	t.Log("fail?")
	actual := math.Abs(a)
	if actual != expect {
		//t.Fatalf("a = %f, actual = %f, expected = %f", a, actual, expect)
		//报告错误
		t.Error("error")
		//t.Fail()
		//t.Fatalf("a = %f, actual = %f, expected = %f", a, actual, expect)
	}
}
