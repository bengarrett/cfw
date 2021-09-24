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

const (
	ellipsis    = "..."
	hexadecimal = 16
)

// DeObfuscate de-obfuscates a CFWheels obfuscateParam or Obfuscate() obfuscated string.
func DeObfuscate(s string) string {
	const twoChrs, decimal = 2, 10
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if _, err := strconv.Atoi(s); err == nil || len(s) < twoChrs {
		return s
	}
	// De-obfuscate string.
	tail := s[twoChrs:]
	n, err := strconv.ParseInt(tail, hexadecimal, 0)

	if err != nil {
		return s
	}

	n ^= 461 // bitxor
	ns := strconv.Itoa(int(n))
	l := len(ns) - 1
	tail = ""

	for i := 0; i < l; i++ {
		f := ns[l-i:][:1]
		tail += f
	}
	// Create checks.
	ct := 0
	l = len(tail)

	for i := 0; i < l; i++ {
		chr := tail[i : i+1]
		n, err1 := strconv.Atoi(chr)

		if err1 != nil {
			return s
		}

		ct += n
	}
	// Run checks.
	ci, err := strconv.ParseInt(s[:2], hexadecimal, 0)
	if err != nil {
		return s
	}

	c2 := strconv.FormatInt(ci, decimal)

	const unknown = 154

	if strconv.FormatInt(int64(ct+unknown), decimal) != c2 {
		return s
	}

	return tail
}

// Excerpt replaces n characters from s which match the first instance of a given phrase.
func Excerpt(s, replace, phrase string, n int) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = ellipsis
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

	mid := ""

	if ln := ep + len(phrase); ln >= len(s) {
		mid = s[sp:]
	} else {
		mid = s[sp:ln]
	}

	return ts + mid + te
}

// Humanize returns readable text by separating camelCase strings to multiple, capitalized words.
func Humanize(s string, except ...string) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/632ea90547da368cddd77cefe17f42a7eda871e0/wheels/global/util.cfm
	//
	// Add a space before every capitalized word.
	s = regexp.MustCompile(`([A-Z])`).ReplaceAllString(s, " $1")
	// Fix abbreviations so they form a word again (example: aURLVariable).
	s = regexp.MustCompile(`([A-Z])\s([A-Z])(?:\s|\b)`).ReplaceAllString(s, "$1$2")
	// Handle exceptions.
	for _, e := range except {
		// (?i) case-insensitive
		s = regexp.MustCompile(`(?i)`+e+`(?:\b)`).ReplaceAllString(s, e)
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
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	//
	// Make sure string doesn't start with a zero.
	if s == "" || s[0] == '0' {
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

	return fmt.Sprintf("%s%s",
		strconv.FormatInt(int64(b), hexadecimal),
		strconv.FormatInt(int64(a), hexadecimal))
}

// StripLinks removes all HTML links from a string leaving just the link text.
func StripLinks(s string) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
	return regexp.MustCompile(`<a.*?>(.*?)</a>`).ReplaceAllString(s, "$1")
}

// StripTags removes all HTML tags from a string.
func StripTags(s string) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
	s = regexp.MustCompile(`<\ *[a-z].*?>`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`<\ */\ *[a-z].*?>`).ReplaceAllString(s, "")

	return s
}

// TimeDistance describes the difference between two time values.
func TimeDistance(from, to time.Time, seconds bool) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	delta := to.Sub(from)
	secs, mins, hrs := int(delta.Seconds()), int(delta.Minutes()), int(delta.Hours())

	var d, m, y int

	const (
		parthour, abouthour, hours, day, days = 45, 90, 1440, 2880, 43200
		month, months, year, years, twoyears  = 86400, 525600, 657000, 919800, 1051200
	)

	switch {
	case mins <= 1:
		if !seconds {
			return lessMin(secs)
		}

		return lessMinAsSec(secs)
	case mins < parthour:
		return fmt.Sprintf("%d minutes", mins)
	case mins < abouthour:
		return "about 1 hour"
	case mins < hours:
		return fmt.Sprintf("about %d hours", hrs)
	case mins < day:
		return "1 day"
	case mins < days:
		const hoursinaday = 24
		d = hrs / hoursinaday

		return fmt.Sprintf("%d days", d)
	case mins < month:
		return "about 1 month"
	case mins < months:
		const hoursinamonth = 730
		m = hrs / hoursinamonth

		return fmt.Sprintf("%d months", m)
	case mins < year:
		return "about 1 year"
	case mins < years:
		return "over 1 year"
	case mins < twoyears:
		return "almost 2 years"
	default:
		y = mins / months

		return fmt.Sprintf("over %d years", y)
	}
}

func lessMin(secs int) string {
	const minute = 60

	switch {
	case secs < minute:
		return "less than a minute"
	default:
		return "1 minute"
	}
}

func lessMinAsSec(secs int) string {
	const five, ten, twenty, forty = 5, 10, 20, 40

	switch {
	case secs < five:
		return "less than 5 seconds"
	case secs < ten:
		return "less than 10 seconds"
	case secs < twenty:
		return "less than 20 seconds"
	case secs < forty:
		return "half a minute"
	default:
		return "1 minute"
	}
}

// Truncate a string to the specified number and replace the trailing characters.
func Truncate(s, replace string, n int) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = ellipsis
	}

	if utf8.RuneCountInString(s) <= n {
		return s
	}

	return s[0:n-utf8.RuneCountInString(replace)] + replace
}

// WordTruncate truncates a string to the specified number of words and replaces the trailing characters.
func WordTruncate(s, replace string, n int) string {
	// CFML source:
	// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = ellipsis
	}

	words := strings.Fields(s)
	if len(words) >= utf8.RuneCountInString(s) {
		return s
	}

	str := ""

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
		return 0, fmt.Errorf("reverseInt %d: %w", i, err)
	}

	return reverse, nil
}
