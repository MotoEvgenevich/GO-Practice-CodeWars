package main

import (
	"fmt"
	"http"
	"ioutil"
	"strings"
)

func main() {
	fmt.Println(MakeUpperCase("hello world!"))
	fmt.Println(Hello())
}

func MakeUpperCase(str string) string {
	sliceLit := []byte{}
	sliceLitUpp := []byte{}
	for _, lit := range str {
		sliceLit = append(sliceLit, byte(lit))
	}
	fmt.Println(sliceLit)

	for _, val := range sliceLit {
		if val >= 97 && val <= 122 {
			sliceLitUpp = append(sliceLitUpp, val-32)
		} else {
			sliceLitUpp = append(sliceLitUpp, val)
		}
		fmt.Println(string(sliceLitUpp))
	}
	return string(sliceLitUpp)
}

func Hello() string {
	slice := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33}
	return string(slice)
}

func greet() string {
	URL := "https://raw.githubusercontent.com/dmfed/go-exercises/master/hello.txt"
	resp, err := http.Get(URL)
	defer resp.Body.Close()
	if err != nil {
		return "Oh, noooo. Could not connect to the internet!"
	}
	out, _ := ioutil.ReadAll(resp.Body)
	return strings.Trim(string(out), "\n")
}
