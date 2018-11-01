package main

import (
	"strings"
	"fmt"
)

func main()  {
	s := "  / The Shawshank Redemption  / 月黑高飞(港) / 刺激1995(台) [可播放]"
	ss := " / American Beauty"

	s1 := strings.TrimLeft(s, "  / ")
	fmt.Printf("trimelect :%s", s1)
}