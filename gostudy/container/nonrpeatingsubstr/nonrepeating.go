package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int) // 定义一个空的字典，没有值的话ok判断为false
	start := 0
	maxlength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1 // 字典中存在想通的key,则移动start 的位置，进行下个的子串的长度统计
		}

		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}

		lastOccurred[ch] = i // 每个字符为key ,索引为值，进行字典的赋值
	}

	return maxlength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkewqazwsxedcrfvtgb"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaaaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("sadfsdgfhd"))
	fmt.Println(lengthOfNonRepeatingSubStr("我不是你你要我怎样"))

}
