package Tool

import (
	"regexp"
)

// RegexpMobile 手机号验证
func RegexpMobile(mobile string) bool {
	ok, err := regexp.MatchString(`^1[3-9][0-9]{9}$`, mobile)
	if err != nil {
		return false
	}
	return ok
}
