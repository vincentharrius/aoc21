package utils

func StrContains(needle string, haystack []string) bool {
	for _, haystraw := range haystack {
		if haystraw == needle {
			return true
		}
	}
	return false
}

func IntContains(needle int, haystack []int) bool {
	for _, haystraw := range haystack {
		if haystraw == needle {
			return true
		}
	}
	return false
}
