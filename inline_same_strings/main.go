package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	groups := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		a := strings.Split(strs[i], "")
		sort.Strings(a)
		key := strings.Join(a, "")
		groups[key] = append(groups[key], strs[i])
	}

	for _,group := range groups {
		fmt.Println(group)
	}

}