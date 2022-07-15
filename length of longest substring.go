package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	length, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		for ; dict[index]; j++ {
			length--
			dict[s[j]] = false
		}

		dict[index] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}
