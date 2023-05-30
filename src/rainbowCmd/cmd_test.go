package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	command := exec.Command("ls")
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func TestIoutil1(t *testing.T) {
	//file, err := ioutil.ReadFile("main.go")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(file)
	dir, err := ioutil.ReadDir("D:/program")
	if err != nil {
		fmt.Println(err)
	}
	println(dir)
}
