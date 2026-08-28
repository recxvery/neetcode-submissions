func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	d1 := make(map[string]int)
	d2 := make(map[string]int)

	for _, v := range s {
		d1[string(v)]++
	}

	for _, v := range t {
		d2[string(v)]++
	}

	for _, v := range s {
		if d1[string(v)] != d2[string(v)] {
			return false
		}
	}

	return true
}
