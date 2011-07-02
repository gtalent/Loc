package main

//Returns a Text type representing the files conents, and a line count
func parseGoFile(val string) (Text) {
	var file Text
	var prev byte
	for i := 0; i < len(val); i++ {
		a := val[i]
		switch a {
		case ' ', '\t', '\n':
			if a == '\n' {
				file.lines++
			}
			file.whitespace++
		case '*', '/':
			if prev == '/' {
				l, chars := parseGoComments(val[i:])
				file.comments += chars
				file.lines += l
				i += chars
				prev = 0
			} else {
				if a == '/' && i + 1 < len(val) && val[i] == '/' {
					for _, a := range val {
						i++
						if a == '\n' {
							file.lines++
							break
						}
					}
				} else {
					file.code++
				}
			}
		default:
			file.code++
		}
		prev = a
	}
	return file
}

func parseGoComments(val string) (int, int) {
	lbs := 0
	chars := 0
	for i := 0; i < len(val); i++ {
		if val[i] == '\n' {
			lbs++
		}
		if val[i] == 42 && i+1 < len(val) && val[i+1] == '/' {
			break
		}
		chars++
	}
	return lbs, chars + 1
}
