package calculator

import "fmt"

func Add(a, b int) int { //ชื่อตัวพิมพ์ใหญ่ คือการทำให้เป็น public
	fmt.Println("Multiply", multiply(2, 3))
	return a + b
}

//fn เวลารับตัวแปรจะเรียกว่า param และ ตอนเรียกใช้ จะเรียนว่า argument
