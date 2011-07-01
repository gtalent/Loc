package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	total := 0
	for _, a := range args {
		file, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Println(err.String())
			continue
		}
		total += strings.Count(string(file), "\n") + 1
	}
	fmt.Println(total, "lines.")
}

