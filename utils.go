package main

import (
	"fmt"
	"os"
)

func throwError(e error) {
	fmt.Println(e.Error())
	os.Exit(0)
}
