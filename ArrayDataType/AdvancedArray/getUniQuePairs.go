package advancedarray

import "fmt"

func GetUniQuePairs() {
	arr := [][2]int{
		{1, 2},
		{3, 2},
		{2, 4},
		{2, 1},
		{5, 6},
		{4, 2},
	}

	seen := make(map[string]bool)
	result := [][2]int{}

	for _, pair := range arr {
		a, b := pair[0], pair[1]
		key1 := fmt.Sprintf("%d-%d", a, b)
		key2 := fmt.Sprintf("%d-%d", b, a)
		if !seen[key1] && !seen[key2] {
			seen[key1] = true
			result = append(result, pair)
		}
	}
	fmt.Println("GetUniQuePairs :", result)
}
