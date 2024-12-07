package variable

import "fmt"

func Variable() {
	//Implicit type  ประกาศไทยตั้งแต่เริ่ม
	var a int = 10

	//Dynamic type ประกาศให้มัน assign type เอง ตาม value
	b := 20

	fmt.Println(a)
	fmt.Println(b)

}
