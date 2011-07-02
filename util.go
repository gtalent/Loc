package main

import "strconv"

type Text struct {
	lines, whitespace, comments, code int
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
	r := ""
	r += "Lines:\t" + strconv.Itoa(me.lines) + "\n"
	r += "Characters:\t" + strconv.Itoa(me.whitespace + me.comments + me.code) + "\n"
	r += "Whitespace:\t%" + strconv.Itoa(me.percentOfFile(me.whitespace)) + "\n"
	r += "Comments:\t%" + strconv.Itoa(me.percentOfFile(me.comments)) + "\n"
	r += "Code:\t\t%" + strconv.Itoa(me.percentOfFile(me.code)) + "\n"
	return r
}
