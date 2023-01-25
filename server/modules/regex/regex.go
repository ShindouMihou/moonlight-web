package regex

import "regexp"

var BearerPrefixRegex = regexp.MustCompile("(?i)(bearer) ")
