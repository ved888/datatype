package advancedarray

import "fmt"

func SearchForAnElementInAnUnsorted(arr [5]int, num int) {
	for _, v := range arr {
		if v == num {
			fmt.Println("this element is found :", num)
			return
		}
	}
	fmt.Println("Element not found:", num)
}

func SearchForAnElementInAnUnsortedString(arr [5]string, text string) {
	for _, v := range arr {
		if v == text {
			fmt.Println("this text is found :", text)
			return
		}
	}
	fmt.Println("Text not found:", text)

}
