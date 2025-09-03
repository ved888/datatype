package question

func RotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	n := len(s)

	for i := 0; i < n; i++ {
		rotate := s[i:] + s[:i]
		if rotate == goal {
			return true
		}
	}
	return false
}
