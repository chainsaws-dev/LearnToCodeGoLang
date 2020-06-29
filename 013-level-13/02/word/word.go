// Package word - contains functions for statistic analysis of texts
package word

import (
	"strings"
)

// UseCount - converts strings to map
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count - counts words in quote
func Count(s string) int {
	return len(strings.Fields(s))
}
