package main

import "fmt"

func main() {
	var x int
	for {
		fmt.Println("enter the number to check:")
		fmt.Scanf("%d", &x)
		if isPalindrome(x) {
			fmt.Println("It's true")
		} else {
			fmt.Println("It's false")
		}
		fmt.Println("-------------")
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	// x = 121
	// 12 1
	// 1 12
	reserveX := 0
	resX := x

	stopCondition := (reserveX == x || resX < 1)
	for {
		if stopCondition {
			break
		}

		tmpResX := resX                                  // 121
		resX = resX / 10                                 // 12
		reserveX = (reserveX * 10) + (tmpResX - resX*10) // 0 * 10 + 121 - 12 * 10 = 0 + 121 - 120 = 1

		stopCondition = (reserveX == x || resX < 1) // 1 == 121 || 12 < 1
	}

	return reserveX == x
}
