package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func fileExt(val string) string {
	ret := ""
	for i := len(val) - 1; i > -1; i-- {
		ret = string(val[i]) + ret
		if val[i] == '.' {
			break
		}
	}
	return ret
}

func buildFileList(path string) (retval []string) {
	return
}

func main() {
	text := make([]*Text, 0)
	langs := func() map[string]*Text {
		retval := make(map[string]*Text)
		retval["go"] = new(Text)
		retval["go"].language = "Go"
		text = append(text, retval["go"])
		return retval
	}()
	var t GenericText
	flag.Parse()
	args := flag.Args()
	for _, a := range args {
		file, err := ioutil.ReadFile(a)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch fileExt(string(a)) {
		case ".go":
			data := parseGoFile(string(file))
			langs["go"].Append(data)
			langs["go"].files++
		default:
			t.Parse(string(file))
		}
	}
	if t.lines > 0 {
		fmt.Println(t.String())
	}
	for _, a := range langs {
		if a.files > 0 {
			fmt.Println(a.String())
		}
	}
}
