package main

// 10. Regular Expression Matching
func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	var dp func(i, j int) bool

	dp = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}
		key := string(i) + "," + string(j)
		if val, found := memo[key]; found {
			return val
		}

		firstMatch := (i < len(s)) && (s[i] == p[j] || p[j] == '.')

		if j+1 < len(p) && p[j+1] == '*' {
			memo[key] = dp(i, j+2) || (firstMatch && dp(i+1, j))
			return memo[key]
		} else if firstMatch {
			memo[key] = dp(i+1, j+1)
			return memo[key]
		}

		memo[key] = false
		return false
	}

	return dp(0, 0)
}
