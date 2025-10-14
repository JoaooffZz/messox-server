package regex

import (
	"regexp"
)

func IsStringInt(id string) bool {
    valid := regexp.MustCompile(`^\d+$`)
    return valid.MatchString(id)
}