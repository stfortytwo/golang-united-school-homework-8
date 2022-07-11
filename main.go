package main

import (
	"io"
	"os"
)

func Perform(args Arguments, writer io.Writer) error {
	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0666)
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
