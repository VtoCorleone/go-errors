package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func readFileError() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", file)
}

func basicError() {
	err := errors.New("Something broke")
	if err != nil {
		fmt.Println(err)
	}
}

func half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func callErrorFromHalf() {
	n, err := half(19)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}

func executePanic() {
	fmt.Println("This is executed")
	panic("Oh no. I can do no more. Goodbye.")
	fmt.Println("This is not executed")
}

func main() {
	readFileError()
	basicError()
	callErrorFromHalf()
	executePanic()
}
