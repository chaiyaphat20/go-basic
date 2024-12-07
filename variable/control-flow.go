package variable

import "fmt"

func Controlflow() {
	//พวก if else switch-case

	n := 3
	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("ODD")
	}

	//short declaration in if statement
	/*if ตัวแปร := เงื่อนไข; ตัวแปร {
		// คำสั่งในกรณีที่เงื่อนไขเป็นจริง
	} else {
		// คำสั่งในกรณีที่เงื่อนไขเป็นเท็จ
	}
	*/

	if a := 10 > 20; a {
		fmt.Println("10 มากกว่า 20")
	} else {
		fmt.Println("10 น้อยกว่า 20")
	}
}
