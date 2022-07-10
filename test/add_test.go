package main

import "testing"

func TestTest(t *testing.T) {
	// 测试数据准备
	data := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},

		{30000, 40000, 50000},
	}

	for _, v := range data {
		if rlt := calcTriangle(v.a, v.b); rlt != v.c {
			t.Errorf("%d 和%d的计算结果应该是%d,得到的是%d", v.a, v.b, v.c, rlt)
		}
	}

}

func TestRepeat(t *testing.T) {
	data := []struct {
		s   string
		ans int
	}{
		{"abc", 3},
		{"sdff", 3},
		{"", 0},
		{"abbsd", 3},
		{"我是", 2},
		{"一二一", 2},
	}

	for _, v := range data {
		if rlt := lengthOfNonRepeatingSubStr(v.s); rlt != v.ans {
			t.Errorf("%s 的结果为%d,应该为 %d", v.s, rlt, v.ans)
		}
	}
}

/**
 * go 性能测试
**/
func BenchmarkSubstr(b *testing.B) {
	s := "一二一二三四为首的发"
	ans := 8

	b.ResetTimer();
	// 进行性能测试的时候，具体计算多少遍系统可以自己决定
	for i := 0; i <= b.N; i++ {
		rlt := lengthOfNonRepeatingSubStr(s)
		if rlt != ans {
			b.Errorf("测试出错，应该是%d,结果是%d", ans, rlt)
		}
	}
}
