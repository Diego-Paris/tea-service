package utils

// AddTrailingSlash adds a trailing slash if the given
// string doesn't already have one.
func AddTrailingSlash(str string) string {

	if len(str) == 0 {
		return "/"
	}

	result := str

	last := string(str[len(str)-1])

	if last != "/" {
		result = result + "/"
	}

	return result
}
