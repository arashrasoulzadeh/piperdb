// Package utils /*
// Copyright © 2025 Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>
package utils

import "regexp"

func IsValidString(s string) bool {
	// Define a regular expression pattern to match:
	// ^[^[:space:][:punct:]\p{So}\p{Cntr}]$ means:
	// ^: start of the string
	// [^...]: not any of the listed characters
	// [:space:]: space characters
	// [:punct:]: punctuation characters
	// \p{So}: Unicode characters categorized as symbols (used for emojis)
	// \p{Cntr}: Unicode control characters
	// $: end of the string
	re := regexp.MustCompile(`^[A-Za-z0-9_]+$`)
	return re.MatchString(s)
}
