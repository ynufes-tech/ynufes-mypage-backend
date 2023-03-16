package testutil

func CheckWithoutOrder[T interface{}](expected, target []T, checkFunc func(T, T) bool) bool {
	if len(expected) != len(target) {
		return false
	}
	for e := range expected {
		for t := range target {
			if checkFunc(expected[e], target[t]) {
				if len(expected) == t-1 {
					target = target[:t]
				} else {
					target = append(target[:t], target[t+1:]...)
				}
				break
			}
		}
	}
	return true
}
