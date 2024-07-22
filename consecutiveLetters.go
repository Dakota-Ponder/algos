package main

import (
	"sort"
	"strings"
)

func Solve(s string) bool {
	sA := strings.Split(s, "")
	sort.Strings(sA)
	s = strings.Join(sA, "")
	return strings.Contains("abcdefghijklmnopqrstuvwxyz", s)
}
