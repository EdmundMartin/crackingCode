package strings_array


func IsUnique(str string) bool {
	mapping := make(map[rune]bool)
	for _, r := range str {
		_, ok := mapping[r]
		if ok {
			return false
		}
		mapping[r] = true
	}
	return true
}
