package slice

import "fmt"

//func main() {
//	var a []int
//	var b [6]int = [6]int{1, 2, 3, 4, 5, 6}
//	//a = b[0:2]
//	//a = b[0:0]
//	//a = b[2:]
//	a = b[0:6]
//	//a = b[:]
//	fmt.Printf("a = %v\n", a)
//}

func main() {
	var sa = make([]string, 5, 10)

	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
		a := fmt.Sprintf("%v", i)
		fmt.Println(a) // [     0 1 2 3 4 5 6 7 8 9]
	}

	for i := 0; i < 5; i++ {
		fmt.Println(sa[i])
	}
	fmt.Println(sa, len(sa))

}
