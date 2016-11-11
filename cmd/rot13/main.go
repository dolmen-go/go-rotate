package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mcandre/go-rotate"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(rotate.Rot13(scanner.Text()))
	}
}
