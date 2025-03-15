package main

import "fmt"

func getValueBySymbol(symbol uint8) int {
	switch symbol {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

// complexity: O(13n)
// 7 characters
// n is len of string
// 7 * n + (7 - 1) * n = 13n

func romanToInt(s string) int {
	if len(s) > 15 {
		return 0
	}

	result := 0
	for i := 0; i < len(s); i++ {
		value1 := getValueBySymbol(s[i])

		if i+1 < len(s) {
			value2 := getValueBySymbol(s[i+1])

			if value2 > value1 {
				result += (value2 - value1)
				i += 1
				continue
			}
		}

		result += value1
	}

	return result
}

func main() {
	var s string
	for {
		fmt.Println("enter the string to convert:")
		fmt.Scanf("%s", &s)

		convertedNumber := romanToInt(s)
		if convertedNumber > 3999 {
			fmt.Println("something was wrong")
		} else {
			fmt.Printf("Result: %d \n", convertedNumber)
		}

		fmt.Println("-------------")
	}
}
