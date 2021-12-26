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

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
