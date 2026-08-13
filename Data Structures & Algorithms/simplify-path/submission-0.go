func simplifyPath(path string) string {
	stack := make([]string, 0)

	for _, p := range strings.Split(path, "/") {
		if p == "" || p == "." {
			continue
		}

		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			} 
		} else {
			stack = append(stack, p)
		}
	}

	return "/" + strings.Join(stack, "/")
}
