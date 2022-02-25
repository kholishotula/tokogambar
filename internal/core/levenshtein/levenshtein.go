package levenshtein

func DistanceTwoStrings(string1, string2 string) int {
	str1rune := []rune(string1)
	str2rune := []rune(string2)

	lenStr1rune := len(str1rune)
	lenStr2rune := len(str2rune)

	if lenStr1rune == 0 {
		return lenStr2rune
	} else if lenStr2rune == 0 {
		return lenStr1rune
	}

	mapColumn := make([][]int, lenStr2rune+1)
	var substitutionCost int

	for i := range mapColumn {
		mapColumn[i] = make([]int, lenStr1rune+1)
		mapColumn[i][0] = i
	}
	for j := 1; j <= lenStr1rune; j++ {
		mapColumn[0][j] = j
	}
	for j := 1; j <= lenStr1rune; j++ {
		for i := 1; i <= lenStr2rune; i++ {
			if str1rune[j-1] == str2rune[i-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			mapColumn[i][j] = min(
				min(mapColumn[i-1][j]+1,
					mapColumn[i][j-1]+1),
				mapColumn[i-1][j-1]+substitutionCost)
		}
	}
	return mapColumn[lenStr2rune][lenStr1rune]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
