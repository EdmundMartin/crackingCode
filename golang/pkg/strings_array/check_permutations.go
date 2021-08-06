package strings_array

import (
	"reflect"
	"sort"
	"strings"
)

func srt(str string) []string  {
	strSlice := strings.Split(str, "")
	sort.Strings(strSlice)
	return strSlice
}


func PermutationWithSorting(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	return reflect.DeepEqual(srt(first), srt(second))
}

func PermutationWithRuneCount(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	mapping := make(map[rune]int)

	for _, r := range first {
		mapping[r]++
	}

	for _, r := range second {
		mapping[r]--
		val, _ := mapping[r]
		if val < 0 {
			return false
		}
	}
	return true
}