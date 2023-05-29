package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for i := 0; i < len(output); i++ {
		r, g, b := rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[i])
	}
}

func main() {
	//var str []string
	//for i := 1; i < 3; i++ {
	//	str = append(str, faker.Hacker().Phrases()...)
	//}
	info, _ := os.Stdin.Stat()
	var out []rune
	fmt.Println(info)
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes")
		fmt.Println("try to print -fortune | go run main.go")
	}
	reader := bufio.NewReader(os.Stdin)
	j := 0
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		out = append(out, input)
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
		j++
	}

}
