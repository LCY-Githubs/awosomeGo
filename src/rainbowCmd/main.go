package main

import (
	"fmt"
	"math"
	"strings"
	"syreclabs.com/go/faker"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func main() {
	var str []string
	for i := 1; i < 3; i++ {
		str = append(str, faker.Hacker().Phrases()...)
	}
	r, g, b := 255, 215, 0 //gold color

	//fmt.Println(strings.Join(str[:], ";"))
	res := strings.Join(str[:], ";")
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		r, g, b = rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, res[i])
	}
}
