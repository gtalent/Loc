package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var output Text
	for _, a := range args {
		file, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Println(err.String())
			continue
		}
		data := parseGoFile(string(file))
		output.Append(data)
	}
	fmt.Println(output.String())
}
