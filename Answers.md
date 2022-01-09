# s-movies
1. simple database querying: 
SELECT a.id, a.name, a.parent as parentId, b.name as parentName FROM member as a LEFT JOIN member as b ON a.parent=b.id

2. movie search: 
this mini project repo

3. refactore code
refactored code:
```golang
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			}
		}
	}
	return ""
}
```

4. anagram
``` golang
func filterAnagram(strArray []string) [][]string {

	var result [][]string
	var tempArr, filteredArr []string
	var tempStr string

	for len(strArray) != 0 {
		for i := 0; i < len(strArray); i++ {
			if tempStr == "" {
				tempStr = strArray[i]
				tempArr = append(tempArr, tempStr)
			} else {
				isAnagram := checkAnagram(tempStr, strArray[i])
				if isAnagram {
					tempArr = append(tempArr, strArray[i])
				} else {
					filteredArr = append(filteredArr, strArray[i])
				}
			}
		}
		result = append(result, tempArr)
		strArray = filteredArr
		filteredArr = nil
		tempArr = nil
		tempStr = ""
	}

	return result
}

func checkAnagram(strA, strB string) bool {

	strArrA := []rune(strA)
	strArrB := []rune(strB)

	if len(strArrA) != len(strArrB) {
		return false
	}

	Sort(strArrA)
	Sort(strArrB)
	return string((strArrA)) == (string(strArrB))
}

func Sort(strArr []rune) {
	sort.Slice(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})
}
```
