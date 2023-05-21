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
