// Package cfw contains ports of a few selected CFWheels helpers that
// are used for string manipulation and have no Go equivalent.
// © Ben Garrett https://github.com/bengarrett/cfw
package cfw

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// DeObfuscate de-obfuscates a CFWheels obfuscateParam or
//  Obfuscate()
// obfuscated string.
func DeObfuscate(s string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if _, err := strconv.Atoi(s); err == nil || len(s) < 2 {
		return s
	}
	// De-obfuscate string.
	new := s[2:] // last 2 chars
	zi, err := strconv.ParseInt(new, 16, 0)
	if err != nil {
		return s
	}
	zi = zi ^ 461 // bitxor
	zs := strconv.Itoa(int(zi))
	l := len(zs) - 1
	new = ""
	for i := 0; i < l; i++ {
		f := zs[l-i:][:1]
		new += f
	}
	// Create checks.
	ct := 0
	l = len(new)
	for i := 0; i < l; i++ {
		chr := new[i : i+1]
		rvi, err := strconv.Atoi(chr)
		if err != nil {
			return s
		}
		ct = ct + rvi
	}
	// Run checks.
	ci, err := strconv.ParseInt(s[:2], 16, 0)
	if err != nil {
		return s
	}
	c2 := strconv.FormatInt(int64(ci), 10)
	if strconv.FormatInt(int64(ct+154), 10) != c2 {
		return s
	}
	return new
}

// Excerpt replaces n characters from s which match the first instance of a given phrase.
func Excerpt(s, replace, phrase string, n int) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = "..."
	}
	pos := strings.Index(s, phrase)
	if pos < 0 {
		return ""
	}
	// Set start info based on whether the excerpt text found, including its radius, comes before the start of the string.
	sp, ts := 0, ""
	if (pos - n) > 1 {
		sp = pos - n
		ts = replace
	}
	// Set end info based on whether the excerpt text found, including its radius, comes after the end of the string.
	ep, te := len(s), ""
	if (pos + len(phrase) + n) <= len(s) {
		ep = pos + n
		te = replace
	}
	var mid string
	if ln := ep + len(phrase); ln >= len(s) {
		mid = s[sp:]
	} else {
		mid = s[sp:ln]
	}
	return ts + mid + te
}

// Humanize returns readable text by separating camelCase strings to multiple, capitalized words.
func Humanize(s, except string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/632ea90547da368cddd77cefe17f42a7eda871e0/wheels/global/util.cfm
	// Add a space before every capitalized word.
	s = regexp.MustCompile(`([A-Z])`).ReplaceAllString(s, " $1")
	// Fix abbreviations so they form a word again (example: aURLVariable).
	s = regexp.MustCompile(`([A-Z])\s([A-Z])(?:\s|\b)`).ReplaceAllString(s, "$1$2")
	// Handle exceptions.
	if except != "" {
		for _, e := range strings.Fields(except) {
			// (?i) case-insensitive
			s = regexp.MustCompile(`(?i)`+e+`(?:\b)`).ReplaceAllString(s, e)
		}
	}
	// Support multiple word input by stripping out all double spaces created.
	s = regexp.MustCompile(`(\s\s)`).ReplaceAllString(s, " ")
	// Capitalize the first letter and trim final result.
	s = strings.TrimPrefix(s, " ")
	return strings.Title(s)
}

// Hyphenize converts camelCase strings to a lowercase hyphened string.
func Hyphenize(s string) string {
	s = regexp.MustCompile(`([A-Z][a-z])`).ReplaceAllString(s, strings.ToLower(`-$1`))
	s = regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllString(s, strings.ToLower(`$1-$2`))
	s = regexp.MustCompile(`^-`).ReplaceAllString(strings.ToLower(s), "")
	return s
}

// Obfuscate a numeric string to insecurely hide database primary key values when passed along a URL.
func Obfuscate(s string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	// Make sure string doesn't start with a zero.
	if len(s) <= 0 || s[0] == '0' {
		return s
	}
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	// Count the number of digits in s.
	count := len(s)
	ri, err := reverseInt(atoi)
	if err != nil {
		return s
	}
	f64 := math.Pow10(count) + float64(ri)
	// Keep a and b as int types.
	a, b := int(f64), 0
	for i := 1; i <= count; i++ {
		// Slice individual digits from s and sum the results.
		ps, err := strconv.Atoi(string(s[count-i]))
		if err != nil {
			return s
		}
		b += ps
	}
	// Base64 conversion.
	a ^= 461
	b += 154
	return strconv.FormatInt(int64(b), 16) + strconv.FormatInt(int64(a), 16)
}

// StripLinks removes all HTML links from a string leaving just the link text.
func StripLinks(s string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
	return regexp.MustCompile(`<a.*?>(.*?)</a>`).ReplaceAllString(s, "$1")
}

// StripTags removes all HTML tags from a string.
func StripTags(s string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
	s = regexp.MustCompile(`<\ *[a-z].*?>`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`<\ */\ *[a-z].*?>`).ReplaceAllString(s, "")
	return s
}

// TimeDistance describes the difference between two time values.
func TimeDistance(from time.Time, to time.Time, seconds bool) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	delta := to.Sub(from)
	s, m, h := int(delta.Seconds()), int(delta.Minutes()), int(delta.Hours())
	var days, months, years int
	switch {
	case m <= 1:
		if seconds != true {
			switch {
			case s < 60:
				return "less than a minute"
			default:
				return "1 minute"
			}
		}
		switch {
		case s < 5:
			return "less than 5 seconds"
		case s < 10:
			return "less than 10 seconds"
		case s < 20:
			return "less than 20 seconds"
		case s < 40:
			return "half a minute"
		default:
			return "1 minute"
		}
	case m < 45:
		return fmt.Sprintf("%d minutes", m)
	case m < 90:
		return "about 1 hour"
	case m < 1440:
		return fmt.Sprintf("about %d hours", h)
	case m < 2880:
		return "1 day"
	case m < 43200:
		days = h / 24
		return fmt.Sprintf("%d days", days)
	case m < 86400:
		return "about 1 month"
	case m < 525600:
		months = h / 730
		return fmt.Sprintf("%d months", months)
	case m < 657000:
		return "about 1 year"
	case m < 919800:
		return "over 1 year"
	case m < 1051200:
		return "almost 2 years"
	default:
		years = m / 525600
		return fmt.Sprintf("over %d years", years)
	}
}

// Truncate a string to the specified number and replace the trailing characters.
func Truncate(s, replace string, n int) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = "..."
	}
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	return s[0:n-utf8.RuneCountInString(replace)] + replace
}

// WordTruncate truncates a string to the specified number of words and replaces the trailing characters.
func WordTruncate(s, replace string, n int) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = "..."
	}
	words := strings.Fields(s)
	if len(words) >= utf8.RuneCountInString(s) {
		return s
	}
	var str string
	for i, w := range words {
		if i+1 >= n {
			str += w
			break
		}
		str += w + " "
	}
	return str + replace
}

// reverseInt reverses an integer.
func reverseInt(i int) (reverse int, err error) {
	// credit: Wade73
	// http://stackoverflow.com/questions/35972561/reverse-int-golang
	itoa, str := strconv.Itoa(i), ""
	for x := len(itoa); x > 0; x-- {
		str += string(itoa[x-1])
	}
	reverse, err = strconv.Atoi(str)
	if err != nil {
		return reverse, err
	}
	return reverse, err
}
