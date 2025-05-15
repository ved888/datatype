package searchword

func FindSubString(str, subString string) bool {

	for i := 0; i <= len(str)-len(subString); i++ {
		match := true
		for j := 0; j < len(subString); j++ {
			if str[i+j] != subString[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false

}
