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
		{"一二一", 4},
	}

	for _, v := range data {
		if rlt := lengthOfNonRepeatingSubStr(v.s); rlt != v.ans {
			t.Errorf("%s 的结果为%d,应该为 %d", v.s, rlt, v.ans)
		}
	}
}
