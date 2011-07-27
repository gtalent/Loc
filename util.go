package main

import (
	"strconv"
	"strings"
)

type GenericText struct {
	lines int
}

func (me *GenericText) Parse(val string) int {
	r := strings.Count(val, "\n") + 1
	me.lines += r
	return r
}

func (me *GenericText) String() string {
	return "Generic text:\n\tLines:" + strconv.Itoa(me.lines) + "\n"
}

type Text struct {
	language                                 string
	files, lines, whitespace, comments, code int
}

func (me *Text) Append(text Text) {
	me.lines += text.lines
	me.whitespace += text.whitespace
	me.comments += text.comments
	me.code += text.code
}

func (me *Text) percentOfFile(val int) int {
	v := float64(val)
	total := float64(me.whitespace) + float64(me.comments) + float64(me.code)
	return int((v * 100) / total)
}

func (me *Text) String() string {
	r := "Language: " + me.language + "\n"
	r += "\tLines:\t" + strconv.Itoa(me.lines) + "\n"
	r += "\tCharacters:\t" + strconv.Itoa(me.whitespace+me.comments+me.code) + "\n"
	r += "\tWhitespace:\t%" + strconv.Itoa(me.percentOfFile(me.whitespace)) + "\n"
	r += "\tComments:\t%" + strconv.Itoa(me.percentOfFile(me.comments)) + "\n"
	r += "\tCode:\t\t%" + strconv.Itoa(me.percentOfFile(me.code)) + "\n"
	return r
}
