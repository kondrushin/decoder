package decoder

func Decode(s string) int {

	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(s)+1; i++ {
		oneDigit := s[i-1 : i]
		twoDigits := s[i-2 : i]
		if oneDigit >= "1" {
			dp[i] += dp[i-1]
		}

		if twoDigits >= "10" && twoDigits <= "26" {
			dp[i] += dp[i-2]
		}

	}

	return dp[len(s)]
}
