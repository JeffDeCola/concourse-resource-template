package main

import (
	"fmt"
	"os"
)

func addthis(a int, b int) (temp int) {
	temp = a + b
	return
	//comment
}

func main() {
	var a = 0
	var b = 1

	a = a + 1
	var sum = addthis(a, b)
	fmt.Println("Hello, everyone. How are you today.  Are you doing well.", sum)
	fmt.Println(os.Getenv("GOPATH"))

}
