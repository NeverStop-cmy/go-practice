package basics

import "fmt"

func RunControlFlowDemo() {
	fmt.Println("\n控制语句演示:")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// switch语句
	switch x {
	case 10:
		fmt.Println("x is 10")
	case 20:
		fmt.Println("x is 20")
	default:
		fmt.Println("x is neither 10 nor 20")
	}

	// for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while循环的Go风格
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	k := 0
	for {
		if k >= 3 {
			break
		}
		fmt.Println(k)
		k++
	}

}
