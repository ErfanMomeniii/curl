package curl

type Stack []string

func (st Stack) Push(s string) {
	var newSt Stack
	newSt = append(newSt, s)
	newSt = append(newSt, st...)
	st = newSt
}

func (st Stack) Pop() {
	st = st[1:]
}

func arrayExist(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
