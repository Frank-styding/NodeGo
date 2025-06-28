package interpreter

func readArray(text string, i *int, limit string, separator string) (list []string) {
	value := ""
	c := get(text, i)
	if c != limit[0] {
		return
	}

	advance(i)
	for {
		c = get(text, i)
		if endLine(text, i) {
			break
		}
		if c == limit[1] {
			list = append(list, value)
			advance(i)
			break
		}
		if c == separator[0] {
			list = append(list, value)
			value = ""
		} else {
			value += getS(text, i)
		}
		advance(i)
	}
	return
}
