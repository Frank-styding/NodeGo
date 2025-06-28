package interpreter

func readString(text string, i *int) string {
	value := ""
	for {
		if endLine(text, i) {
			break
		}
		c := get(text, i)
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' {
			value += getS(text, i)
		} else {
			break
		}
		advance(i)
	}
	return value
}
