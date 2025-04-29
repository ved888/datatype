package advancedarray

import (
	"fmt"
	"math"
)

func FindMaximumValue(arr [5]int) {
	max := 0

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("maximum value from an array :", max)
}

func FindMinimumValue(arr [5]int) {
	min := arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	fmt.Println("print minimum value from an array :", min)
}

func FindSecondMaxNumber(arr [5]int) {
	max1, max2 := arr[0], arr[0]

	for _, v := range arr {
		if max1 < v {
			max2 = max1
			max1 = v
		} else if max2 < v && max1 != v {
			max2 = v
		}
	}
	fmt.Println("print second larg number :", max2)
}

func FindSecondMinNumber(arr [5]int) {
	fMin, sMin := arr[0], math.MaxInt

	for _, v := range arr {
		if v < fMin {
			sMin = fMin
			fMin = v
		} else if v < sMin && v != fMin {
			sMin = v
		}
	}
	fmt.Println("second minimum number :", sMin)

}

func FindThirdMaxNumber(arr [10]int) {
	// f, s, t := math.MinInt, math.MinInt, math.MinInt
	f, s, t := arr[0], arr[0], arr[0] // both are currect

	for _, v := range arr {
		if f < v {
			t = s
			s = f
			f = v
		} else if s < v && f != v {
			t = s
			s = v
		} else if t < v && s != v && t != v {
			t = v
		}
	}
	fmt.Println("third largest number :", t)
}

func FindThirdMinNumber(arr [10]int) {
	f, s, t := math.MaxInt, math.MaxInt, math.MaxInt

	for _, v := range arr {
		if f > v {
			t = s
			s = f
			f = v
		} else if s > v && f != v {
			t = s
			s = v
		} else if t > v && f != v && s != v {
			t = v
		}
	}
	fmt.Println("third minimumn number :", t)

}
