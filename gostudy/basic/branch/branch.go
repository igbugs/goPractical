package main

import (
	"fmt"
)

//func main()  {
//	const filename  = "abc.txt"
//	contents, err := ioutil.ReadFile(filename)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%s\n", contents)
//	}
//    fmt.Printf("%s\n",contents)
//}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		g = "default"
	}
	return g
}

func main() {
	//const filename  = "abc.txt"
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	fmt.Println(
		grade(0),
		grade(55),
		grade(67),
		grade(87),
		grade(90),
		grade(100),
		grade(102),
	)
}
