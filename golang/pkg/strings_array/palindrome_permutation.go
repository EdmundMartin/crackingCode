package strings_array


func buildFreqTable(input string) map[rune]int {
	freqTable := make(map[rune]int)
	for _, r := range input {
		freqTable[r]++
	}
	return freqTable
}

func checkMaxOneOdd(freqTable map[rune]int) bool {
	foundOdd := false
	for _, value := range freqTable {
		if value % 2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

func IsPermutationPalindromeNaive(phrase string) bool {
	freqTable := buildFreqTable(phrase)
	return checkMaxOneOdd(freqTable)
}

func IsPermutationPalindromeOptimised(phrase string) bool {
	countOdd := 0
	freqTable := make(map[rune]int)
	for _, r := range phrase {
		freqTable[r]++
		val, _ := freqTable[r]
		if val % 2 == 1 {
			countOdd++
		} else {
			countOdd--
		}
	}
	return countOdd <= 1
}