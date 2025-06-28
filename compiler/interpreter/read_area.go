package interpreter

func readArea(text string, i *int, limit string) (value string) {
	value = ""
	c := get(text, i)
	for get(text, i) == '\n' {
		advance(i)
	}
	if c != limit[0] {
		return
	}
	advance(i)
	for {
		c = get(text, i)
		if endLine(text, i) {
			return
		}
		if c == limit[1] {
			/* advance(i) */
			return
		}
		value += getS(text, i)
		advance(i)
	}
}
