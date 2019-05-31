package utils

import "regexp"

func Includes(a []string, t string) bool {
	for _, s := range a {
		if s == t {
			return true
		}
	}
	return false
}

func Excludes(a []string, t string) (ra []string) {
	for _, s := range a {
		if s != t {
			ra = append(ra, s)
		}
	}
	return
}

func Select(a []string, r *regexp.Regexp) (ra []string) {
	for _, s := range a {
		if r.MatchString(s) {
			ra = append(ra, s)
		}
	}
	return
}