package limiter

func determineSpeedLimit(limit1, limit2 int) int {
	if limit1 == 0 {
		return limit2
	}
	if limit2 == 0 {
		return limit1
	}
	if limit1 < limit2 {
		return limit1
	}
	return limit2
}
