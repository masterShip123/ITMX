package utilities

func FormatSeperatedString(delimiter string, str []string, isPreserveZeroValue bool) string {
	result := ""
	comma := ""
	for _, v := range str {
		if isPreserveZeroValue || v != "" {
			result += comma + v
			comma = delimiter
		}
	}
	return result
}
