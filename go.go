/*
   Copyright 2011-2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package main

//Returns a Text type representing the files conents, and a line count
func parseGoFile(val string) Text {
	var file Text
	for i := 0; i < len(val); i++ {
		a := val[i]
		switch a {
		case ' ', '\t', '\n':
			if a == '\n' {
				file.lines++
			}
			file.whitespace++
		case '/':
			//new
			if !(i+1 < len(val)) { // if does not have next
				file.code++
				break
			}
			i++
			switch val[i] {
			case '/': //starts one line comment
				{
					c := val[i]
					for {
						file.comments++
						if c == '\n' {
							file.lines++
							break
						}
						i++
						if !(i < len(val)) {
							break
						}
						c = val[i]
					}
				}
			case '*': //start comment block
				i++
				lines, chars := parseGoCommentBlock(val[i:])
				file.comments += chars + 2 //the + 2 is for the two uncount "i++"s
				file.lines += lines
				i += chars
			default: //does not start comment
				file.code++
			}
		default:
			file.code++
		}
	}
	return file
}

//Starts processing at the start of comment block's conetents
func parseGoCommentBlock(val string) (int, int) {
	lbs := 0
	chars := 0
	for i := 0; i < len(val); i++ {
		if val[i] == '\n' {
			lbs++
		}
		if val[i] == '*' && i+1 < len(val) && val[i+1] == '/' {
			break
		}
		chars++
	}
	return lbs, chars
}
