package curl

type Stack []string

func PushStack(st Stack, s string) Stack {
	var newSt Stack
	newSt = append(newSt, s)
	newSt = append(newSt, st...)
	return newSt
}

func PopStack(st Stack) Stack {
	if len(st) > 0 {
		return st[1:]
	}
	return nil
}

func arrayExist(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
