package main

import (
	advancedarray "ArrayDataType/AdvancedArray"
	arrayinitialization "ArrayDataType/ArrayInitialization"
	basicarray "ArrayDataType/basicArray"
	"fmt"
)

func main() {
	arr := [5]int{2, 31, 22, 11, 10}
	arr1 := [5]int{2, 33, 27, 11, 10}

	arrayinitialization.ArrayFullInitialization()
	arrayinitialization.PartialInitialization()
	arrayinitialization.IndexBasedInitialization()
	arrayinitialization.CompilerInferredLength()
	arrayinitialization.AccessElements()
	basicarray.CopyArray(arr)
	basicarray.MultiDimensionalArray()
	fmt.Println("return type array :", advancedarray.ReturnAnArray())
	nums := [3]int{1, 2, 3}
	advancedarray.ModifyArray(&nums)
	fmt.Println(nums) // Output: [99 2 3]
	advancedarray.ReverseAnArray(&arr)
	fmt.Println("reverse the array :", arr)
	advancedarray.SortAnArrayInAscending(&arr)
	fmt.Println("sort of the array ascending :", arr)
	advancedarray.SortAnArrayInDescending(&arr1)
	fmt.Println("sort of thr array descending :", arr1)
	arr2 := [5]int{2, 1, 2, 2, 7}
	x := advancedarray.RemoveDuplicatesFromArray(arr2)
	fmt.Println("remove duplicate ilement from array:", x)
	advancedarray.RemoveDuplicatesFromArrayWithoutReturn(&arr2)
	fmt.Println("remove duplicate ilement from array :", arr2)
	mearge := advancedarray.MergeArrays(arr1, arr2)
	fmt.Println("mearge two array :", mearge)
	advancedarray.FindMaximumValue(arr)
	advancedarray.FindMinimumValue(arr)
	arr3 := [5]int{2, 33, 27, 128, 10}
	advancedarray.FindSecondMaxNumber(arr3)
	arr4 := [5]int{2, 33, 27, 128, 10}
	advancedarray.FindSecondMinNumber(arr4)
	arr5 := [10]int{2, 33, 56, 128, 10, 22, 4, 3, 90, 25}
	advancedarray.FindThirdMaxNumber(arr5)
	advancedarray.FindThirdMinNumber(arr5)
	arr6 := [5]int{2, 1, 3, 4, 5}
	advancedarray.SumOfAllElements(arr6)
	arr7 := [5]int{2, 1, 3, 4, 5}
	advancedarray.ShiftLeft(&arr7)
	fmt.Println("sift array left :", arr7)
	advancedarray.ShiftRight(&arr6)
	fmt.Println("sift array right :", arr6)
	advancedarray.SwapTwoElements(&arr7, 1, 3)
	fmt.Println("swap two element in array :", arr7)
	advancedarray.SearchForAnElementInAnUnsorted(arr7, 6)
	text := [5]string{"hello", "q", "world", "jfhj", "sdh"}
	advancedarray.SearchForAnElementInAnUnsortedString(text, "hello")
	slice := []int{5, 2, 9, 2, 7, 5, 3, 3, 9, 6}
	advancedarray.SortBasicOnDuplicateElement(slice)
	fmt.Println("print sort with duplicate :", slice)
	slic := advancedarray.ArrayToSlice(arr7)
	fmt.Println("convert array to slice :", slic)
	names := []string{"Zoe", "Alice", "Bob"}
	advancedarray.SortStringArray(names)
	fmt.Println("sort the string :", names)
	str := advancedarray.SwapPara("Hello world, this is Go")
	fmt.Println("word revers in para is :", str)
	word := advancedarray.ReverseWordOrder("Hello world, this is Go")
	fmt.Println("Reverse the words :", word)
	match := advancedarray.SearchWordFromText("Hello world, this is Go", "Hello")
	fmt.Println("word is match :", match)
	bin := []int{2, 4, 6, 8, 9}
	binary := advancedarray.BinarySearch(bin, 6)
	fmt.Println("this target number is index :", binary)
	sortArr1 := []int{26, 38, 60, 87, 90, 95}
	sortArr2 := []int{13, 35, 50, 59, 85, 89, 91, 98, 99}
	meargArray := advancedarray.MeargTwoSortedArray(sortArr1, sortArr2)
	fmt.Println("MeargTwoSortedArray :", meargArray)
	advancedarray.GetUniQuePairs()
}
