package kata

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}

	n := len(strng)
	for i := 0; i < n; i++ {
		rotation := strng[i:] + strng[:i]
		found := false
		for _, v := range arr {
			if v == rotation {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
