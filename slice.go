package goopt

// Here we have some utility slice routines

func any(f func(string) bool, slice []string) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}
