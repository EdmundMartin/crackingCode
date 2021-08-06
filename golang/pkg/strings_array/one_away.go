package strings_array

import "strings"

func OneAway(first, second string) bool {
	if len(first) == len(second) {
		return oneEditReplace(first, second)
	} else if len(first) - 1 == len(second) {
		return oneEditInsert(first, second)
	} else if len(first) + 1 == len(second) {
		return oneEditInsert(second, first)
	}
	return false
}

func oneEditReplace(first, second string) bool {
	one := strings.Split(first, "")
	two := strings.Split(second, "")
	foundDifference := false
	for idx, char := range one {
		if char != two[idx] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}

func oneEditInsert(first, second string) bool {
	firstIdx, secondIdx := 0, 0
	one := strings.Split(first, "")
	two := strings.Split(second, "")
	for secondIdx < len(second) && firstIdx < len(one) {
		if one[firstIdx] != two[secondIdx] {
			if firstIdx != secondIdx {
				return false
			}
			secondIdx++
		} else {
			firstIdx++
			secondIdx++
		}
	}
	return true
}