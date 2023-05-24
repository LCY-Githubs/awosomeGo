package main

import (
	"fmt"
	"testing"
)

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

//必须以Test开头
//可选的后缀名也必须以大写字母开头
func TestFunc(t *testing.T) {
	//println(IsPalindrome("sdasdasdasdada"))

}

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`Ispalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`Ispalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	input := "palind"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func TestAnonymousFun(t *testing.T) {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func TestAppendFunc1(t *testing.T) {
	//第一种用法
	//func append(slice []Type, elems ...Type) []Type
	nums := []int{1, 2}
	fmt.Println(nums)
	nums = append(nums, 4)
	fmt.Println(nums)
	nums = append(nums, 5, 6, 7)
	fmt.Println(nums)

	//第二种用法
	//res := append(slice1, slice2...)
	//第三种用法
	//append([]byte(string1),string2...)
}
func TestAppendCapBig(t *testing.T) {
	var nums2 []int
	nums2 = make([]int, 2, 2)
	nums2[0] = 0
	nums2[1] = 1
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
	nums2 = append(nums2, 3)
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
}
