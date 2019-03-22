package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dict:= map[interface{}]interface{} {
		1: "hello",
		2: "hey",
	}
	fmt.Print(dict)
	id  := dict[1].(string)
	fmt.Print(id)

	dict[3] = "hi";
	fmt.Print(dict)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)





}
