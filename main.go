package main

import (
	"fmt"

	fo "github.com/NViktorovich/ktssaver/ProgramFiles/FileOperation"
)

func main() {
	path, _ := fo.GetAbsPath()
	fmt.Println(path)

}
