

type Stack struct {
	items []string
}

func isValid(s string) bool {
	pair := map[string]string{
		"{" : "}",
		"(" : ")",
		"[" : "]",
	}

	stack := Stack{
		items: make([]string, 0),
	}

	for _, value := range s {
		str := string(value)
		if _, ok := pair[str]; ok {
			stack.items = append(stack.items, str)
		} else if len(stack.items) == 0 {
			return false 
		} else {
			popstr, IsPopped := stack.Pop()
			if !IsPopped {
				return false
			}

			if pair[popstr] != str {
				return false
			}
		}

	}
	return len(stack.items) == 0
}

func (s *Stack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	
	popped := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]

	return popped, true
}
