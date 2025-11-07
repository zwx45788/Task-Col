package common

import "regexp"

func VerifyMobilde(mobile string) bool {
	if mobile == "" {
		return false
	}
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(mobile)
}
