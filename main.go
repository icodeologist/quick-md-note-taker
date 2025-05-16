package main

import (
	"fmt"
	"github.com/icodeologist/project-chaos/chaos"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	if len(os.Args) < 3 || os.Args[1] != "new" {
		fmt.Println("Usage : chaos new \"your title here\"")
		return
	}
	title := os.Args[2]
	err := chaos.CreateEntry(title)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Entry created successfully!")
}
