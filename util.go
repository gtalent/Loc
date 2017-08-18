/*
   Copyright 2011-2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

import (
	"strconv"
	"strings"
)

type genericText struct {
	lines int
}

func (me *genericText) Parse(val string) int {
	r := strings.Count(val, "\n") + 1
	me.lines += r
	return r
}

func (me *genericText) String() string {
	return "Generic text:\n\tLines:" + strconv.Itoa(me.lines) + "\n"
}

type text struct {
	language                                 string
	files, lines, whitespace, comments, code int
}

func (me *text) Append(text text) {
	me.lines += text.lines
	me.whitespace += text.whitespace
	me.comments += text.comments
	me.code += text.code
}

func (me *text) percentOfFile(val int) int {
	v := float64(val)
	total := float64(me.whitespace) + float64(me.comments) + float64(me.code)
	return int((v * 100) / total)
}

func (me *text) String() string {
	r := "Language: " + me.language + "\n"
	r += "\tLines:\t" + strconv.Itoa(me.lines) + "\n"
	r += "\tCharacters:\t" + strconv.Itoa(me.whitespace+me.comments+me.code) + "\n"
	r += "\tWhitespace:\t%" + strconv.Itoa(me.percentOfFile(me.whitespace)) + "\n"
	r += "\tComments:\t%" + strconv.Itoa(me.percentOfFile(me.comments)) + "\n"
	r += "\tCode:\t\t%" + strconv.Itoa(me.percentOfFile(me.code)) + "\n"
	return r
}
