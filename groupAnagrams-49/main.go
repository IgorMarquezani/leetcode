package main

import (
	"fmt"
	"slices"
)

func groupAnagrams2(strs []string) [][]string {
    m := make(map[[26]uint8][]string)
    
    for i := range strs {
        a26 := toA26(strs[i])
        m[a26] = append(m[a26], strs[i])
    }

    res := make([][]string, 0, len(m))

    for _, sl := range m {
        res = append(res, sl)
    }

    return res
}

func toA26(s string) [26]uint8 {
    var chars [26]uint8

    for i := range s {
        chars[s[i]-'a']++
    }

    return chars
}

func groupAnagrams(strs []string) [][]string {
	var (
		result       = [][]string{}
		combinations = map[string][]string{}
	)

	for i := range strs {
		sorted := []byte(strs[i])
		slices.Sort(sorted)

		val, ok := combinations[string(sorted)]
		if !ok {
			combinations[string(sorted)] = []string{strs[i]}
		} else {
			combinations[string(sorted)] = append(val, strs[i])
		}
	}

	for _, v := range combinations {
		result = append(result, v)
	}

	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r := groupAnagrams(strs)
	fmt.Println(r)
}
