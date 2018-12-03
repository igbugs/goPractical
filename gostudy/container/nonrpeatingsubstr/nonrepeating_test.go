package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s      string
		answer int
	}{
		// Normal cases
		{"abcabcaa", 3},
		{"sadfsdgfhd", 5},
		{"pwwkewqazwsxedcrfvtgb", 15},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"aaaaaa", 1},
		{"abcabcabcd", 4},

		// Chinase support
		{"我不是你你要我怎样", 5},
		{"我在学go语言", 7},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞挥发会灰化肥灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.answer {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.answer)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞挥发会灰化肥灰化肥挥发发黑会飞花"
	answer := 8

	for i := 0; i < b.N; i++ { // b.N 是有算法去决定执行的次数
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != answer {
			b.Errorf("got %d for input %s; expected %d",
				actual, s, answer)
		}
	}
}
