package loop

import "fmt"

func Loop() {
	var sum int

	for i := 1; i <= 10; i++ {
		sum += 1
	}

	for i := range 3 { //[0,1,2]
		fmt.Println(i)
	}
}
