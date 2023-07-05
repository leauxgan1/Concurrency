package main

import (
	"strings"
)

func delimit(text string, delimiter string) string {
	var builder strings.Builder
	var i int = 0
	for i < len(text) {
		char := text[i]
		builder.WriteString(string(char))
		builder.WriteString(delimiter)
		i += 1
	}
	return builder.String()
}

func leftTab(text string, tabs int, tabChar string) string {
	var builder strings.Builder
	i := 0
	for i < tabs {
		i += 1
		builder.WriteString(tabChar)
	}
	builder.WriteString(text)
	return builder.String()
}
