package main

import (
	"fmt"
)

func main() {
	num := 75
	// switch { // 表达式被省略了
	// case num >= 0 && num <= 50:
	// 	fmt.Println("num is greater than 0 and less than 50")
	// case num >= 51 && num <= 100:
	// 	fmt.Println("num is greater than 51 and less than 100")
	// case num >= 101:
	// 	fmt.Println("num is greater than 100")
	// }
	switch { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

}