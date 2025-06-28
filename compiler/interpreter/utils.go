package interpreter

func get(text string, i *int) byte {
	if *i >= len(text) {
		return 0
	}
	return text[*i]
}

func getS(text string, i *int) string {
	if *i >= len(text) {
		return ""
	}
	return string(text[*i])
}

func endLine(text string, i *int) bool {
	return *i >= len(text)
}

func advance(i *int) {
	(*i)++
}

func back(i *int) {
	(*i)--
}
