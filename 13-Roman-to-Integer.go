package main

import "fmt"

// func main() {
// 	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
// 	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// 	myinput := "LVIII"
// 	result := 0

// 	for i := 0; i < len(myinput); i++ {
// 		if i+1 < len(myinput) {
// 			twoChar := string(myinput[i]) + string(myinput[i+1])
// 			found := false
// 			for j := 0; j < len(symbols); j++ {
// 				if symbols[j] == twoChar {
// 					result += values[j]
// 					i++
// 					found = true
// 					break
// 				}
// 			}
// 			if found {
// 				continue
// 			}
// 		}

// 		currentChar := string(myinput[i])
// 		for j := 0; j < len(symbols); j++ {
// 			if symbols[j] == currentChar {
// 				result += values[j]
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(result)
// }

func RomanToInteger() {

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	s := "MCMXCIV"
	res := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if romanMap[s[i]] < romanMap[s[i+1]] {
				res -= romanMap[s[i]]
			} else {
				res += romanMap[s[i]]
			}
		} else {
			res += romanMap[s[i]]
		}
	}
	fmt.Println(res)
}
