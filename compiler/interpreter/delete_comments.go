package interpreter

func deleteComments(text string) (nText string) {
	for i := 0; i < len(text); i++ {
		if get(text, &i) == '#' {
			for !endLine(text, &i) && get(text, &i) != '\n' {
				advance(&i)
			}
			if !endLine(text, &i) && get(text, &i) == '\n' {
				advance(&i)
			}
		}
		if !endLine(text, &i) {
			nText += getS(text, &i)
		}
	}
	return
}
