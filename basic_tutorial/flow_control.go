package basic_tutorial

import "fmt"

func IfControl(x int) string {

	if x > 5 {
		return fmt.Sprintf("%d es mayor a 5", x)
	} else {
		return fmt.Sprintf("%d es menor a 5", x)
	}
}

func ForControl(untilValue int) int {
	total := 0
	for i := 0; i < untilValue+1; i++ {
		fmt.Println(i)
		total += i
	}
	return total
}
