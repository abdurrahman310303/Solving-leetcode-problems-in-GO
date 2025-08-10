package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	keypadMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	result := []string{""}

	for i := 0; i < len(digits); i++ {
		var current []string
		digit := string(digits[i])
		letters := keypadMap[digit]

		for _, combination := range result {
			for j := 0; j < len(letters); j++ {
				current = append(current, combination+string(letters[j]))
			}
		}
		result = current
	}

	return result
}

func letterCombinationsMain() {
	digits := "23"
	result := letterCombinations(digits)
	fmt.Println(result)
}
