package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

}
