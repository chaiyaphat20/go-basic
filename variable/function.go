package variable

import "fmt"

// Non-Pure:
func greet(name string) {
	fmt.Println("Hello, " + name)
}

// Pure:
func greet2(name string) string {
	return "Hello, " + name
}

func Functions() {

	/*
		Pure Function คือฟังก์ชันที่มีคุณสมบัติสำคัญ 2 อย่าง:
		ผลลัพธ์เดียวกันเสมอ: เมื่อเรียกฟังก์ชันด้วยอินพุตเดิม จะได้ผลลัพธ์เดิมเสมอโดยไม่ขึ้นกับสถานะภายนอก
		ไม่มี Side Effects: ฟังก์ชันไม่เปลี่ยนแปลงสถานะใด ๆ ภายนอกตัวฟังก์ชัน เช่น ไม่แก้ไขตัวแปรภายนอก, ไม่เขียนไฟล์, หรือไม่พิมพ์ข้อมูลลงหน้าจอ

	*/

	/*
		Non-Pure Function มีการเปลี่ยนแปลงค่าข้างนอก
		var counter int = 0

		func increment() {
		    counter++ // แก้ไขตัวแปรภายนอก
		}
	*/

	greet("art")
	fmt.Println(greet2("Alice"))

}
