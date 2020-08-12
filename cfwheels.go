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

const ellipsis = "..."

// DeObfuscate de-obfuscates a CFWheels obfuscateParam or a string mutated by
// Obfuscate().
func DeObfuscate(s string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if _, err := strconv.Atoi(s); err == nil || len(s) < 2 {
		return s
	}
	// De-obfuscate string.
	tail := s[2:] // last 2 chars
	zi, err := strconv.ParseInt(tail, 16, 0)
	if err != nil {
		return s
	}
	zi ^= 461 // bitxor
	zs := strconv.Itoa(int(zi))
	l := len(zs) - 1
	tail = ""
	for i := 0; i < l; i++ {
		f := zs[l-i:][:1]
		tail += f
	}
	// Create checks.
	ct := 0
	l = len(tail)
	for i := 0; i < l; i++ {
		chr := tail[i : i+1]
		rvi, errl := strconv.Atoi(chr)
		if errl != nil {
			return s
		}
		ct += rvi
	}
	// Run checks.
	ci, err := strconv.ParseInt(s[:2], 16, 0)
	if err != nil {
		return s
	}
	c2 := strconv.FormatInt(ci, 10)
	const unknown = 154
	if strconv.FormatInt(int64(ct+unknown), 10) != c2 {
		return s
	}
	return tail
}

// Excerpt replaces n characters from s which match the first instance of a given phrase.
func Excerpt(s, replace, phrase string, n int) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
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
	var mid string
	if ln := ep + len(phrase); ln >= len(s) {
		mid = s[sp:]
	} else {
		mid = s[sp:ln]
	}
	return ts + mid + te
}

// Humanize returns readable text by separating camelCase strings to multiple, capitalized words.
func Humanize(s string, except ...string) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/632ea90547da368cddd77cefe17f42a7eda871e0/wheels/global/util.cfm
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
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
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
func TimeDistance(from, to time.Time, seconds bool) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	delta := to.Sub(from)
	sec, mnt, hrs := int(delta.Seconds()), int(delta.Minutes()), int(delta.Hours())
	var d, m, y int
	const (
		min, five, ten, twenty, forty         = 60, 5, 10, 20, 40
		parthour, abouthour, hours, day, days = 45, 90, 1440, 2880, 43200
		month, months, year, years, twoyears  = 86400, 525600, 657000, 919800, 1051200
	)
	switch {
	case mnt <= 1:
		if !seconds {
			switch {
			case sec < min:
				return "less than a minute"
			default:
				return "1 minute"
			}
		}
		switch {
		case sec < five:
			return "less than 5 seconds"
		case sec < ten:
			return "less than 10 seconds"
		case sec < twenty:
			return "less than 20 seconds"
		case sec < forty:
			return "half a minute"
		default:
			return "1 minute"
		}
	case mnt < parthour:
		return fmt.Sprintf("%d minutes", mnt)
	case mnt < abouthour:
		return "about 1 hour"
	case mnt < hours:
		return fmt.Sprintf("about %d hours", hrs)
	case mnt < day:
		return "1 day"
	case mnt < days:
		const hoursinaday = 24
		d = hrs / hoursinaday
		return fmt.Sprintf("%d days", d)
	case mnt < month:
		return "about 1 month"
	case mnt < months:
		const hoursinamonth = 730
		m = hrs / hoursinamonth
		return fmt.Sprintf("%d months", m)
	case mnt < year:
		return "about 1 year"
	case mnt < years:
		return "over 1 year"
	case mnt < twoyears:
		return "almost 2 years"
	default:
		y = mnt / months
		return fmt.Sprintf("over %d years", y)
	}
}

// Truncate a string to the specified number and replace the trailing characters.
func Truncate(s, replace string, n int) string {
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
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
	// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
	if replace == "" {
		replace = ellipsis
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
		return 0, fmt.Errorf("reverseInt %d: %w", i, err)
	}
	return reverse, nil
}
