package main

func in2post(regex string) string {
	result := ""
	charactersStack := make([]byte, 0)
	for _, c := range []byte(regex) {
		if c == '.' || c == '(' {
			charactersStack = append(charactersStack, c)
		} else if c == '|' {
			n := len(charactersStack)
			for n > 0 && charactersStack[n-1] == '.' {
				result = result + string([]byte{charactersStack[n-1]})
				charactersStack = charactersStack[:n-1]
				n -= 1

			}
			charactersStack = append(charactersStack, c)
		} else if c == ')' {
			n := len(charactersStack)
			for n > 0 && charactersStack[n-1] != '(' {
				result = result + string([]byte{charactersStack[n-1]})
				charactersStack = charactersStack[:n-1]
				n -= 1
			}
			if n > 0 && charactersStack[n-1] == '(' {
				charactersStack = charactersStack[:n-1]
			}
		} else {
			result = result + string([]byte{c})
		}
	}
	n := len(charactersStack)
	for n > 0 {
		result = result + string([]byte{charactersStack[n-1]})
		charactersStack = charactersStack[:n-1]
		n -= 1
	}
	return result
}
