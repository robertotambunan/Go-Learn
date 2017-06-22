package main

import "fmt"
import "github.com/user/stringfunction"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Print("Reverse The Word: ")
	fmt.Println(stringfunction.Reverse("Hello World"))
}
