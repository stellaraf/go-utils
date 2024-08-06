package format

import (
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TitleCase converts any string to title case. For example, "UNKNOWN" becomes "Unknown"
func TitleCase(in string) string {
	p := regexp.MustCompile(`\_|\-`)
	in = p.ReplaceAllString(in, " ")
	return cases.Title(language.AmericanEnglish).String(in)
}
