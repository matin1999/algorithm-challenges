package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	fmt.Println("******")
	fmt.Println(anagrams)
	fmt.Println("******")

	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := groupAnagrams(strs)

	for _, group := range grouped {
		fmt.Println(strings.Join(group, " "))
	}
}
