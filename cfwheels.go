// Package cfw contains ports of a few selected CFWheels helpers that
// are used for string manipulation and have no Go equivalent.
// © Ben Garrett
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

// DeObfuscate de-obfuscates s that has previously been obfuscated with Obfuscate().
// A port of CFWheels deobfuscateParam() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func DeObfuscate(s string) string {
	_, err := strconv.Atoi(s)
	if err == nil || len(s) < 2 {
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
	c1 := strconv.FormatInt(int64(ct+154), 10)
	cs := s[:2] // first 2 chars of s
	ci, err := strconv.ParseInt(cs, 16, 0)
	if err != nil {
		return s
	}
	c2 := strconv.FormatInt(int64(ci), 10)
	if c1 != c2 {
		return s
	}
	return new
}

// Excerpt extracts n characters from s which matches the first instance of a given phrase and surrounds the excerpt with new.
// A port of CFWheels excerpt() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func Excerpt(s string, new string, phrase string, n int) string {
	if new == "" {
		new = "..."
	}
	pos := strings.Index(s, phrase)
	// Return an empty value if the text wasn't found at all.
	if pos < 0 {
		return ""
	}
	// Set start info based on whether the excerpt text found, including its radius, comes before the start of the string.
	sp := 0
	ts := ""
	if (pos - n) > 1 {
		sp = pos - n
		ts = new
	}
	// Set end info based on whether the excerpt text found, including its radius, comes after the end of the string.
	ep := len(s)
	te := ""
	if (pos + len(phrase) + n) <= len(s) {
		ep = pos + n
		te = new
	}
	ln := ep + len(phrase)
	var mid string
	if ln >= len(s) {
		mid = s[sp:]
	} else {
		mid = s[sp:ln]
	}
	return ts + mid + te
}

// Humanize returns readable text by capitalizing and converting camel casing of s into multiple words.
// except is an optional list of space separated words to replace within the returned string.
// A port of CFWheels humanize() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/632ea90547da368cddd77cefe17f42a7eda871e0/wheels/global/util.cfm
func Humanize(s string, except string) string {
	// Add a space before every capitalized word.
	re := regexp.MustCompile(`([A-Z])`) // parse regular expression
	s = re.ReplaceAllString(s, " $1")   // run reg-ex and save results to t
	// Fix abbreviations so they form a word again (example: aURLVariable).
	re = regexp.MustCompile(`([A-Z])\s([A-Z])(?:\s|\b)`)
	s = re.ReplaceAllString(s, "$1$2")
	// Handle exceptions.
	if except != "" {
		es := strings.Fields(except)
		for _, e := range es {
			re = regexp.MustCompile(`(?i)` + e + `(?:\b)`) // (?i) case-insensitive
			s = re.ReplaceAllString(s, e)
		}
	}
	// Support multiple word input by stripping out all double spaces created.
	re = regexp.MustCompile(`(\s\s)`)
	s = re.ReplaceAllString(s, " ")
	// Capitalize the first letter and trim final result.
	s = strings.TrimPrefix(s, " ")
	return strings.Title(s)
}

// Hyphenize converts camelCase strings to lowercase strings with hyphens as word delimiters instead.
// A port of CFWheels hyphenize() helper.
// Example: myVariable becomes my-variable.
func Hyphenize(s string) string {
	re := regexp.MustCompile(`([A-Z][a-z])`)
	s = re.ReplaceAllString(s, strings.ToLower(`-$1`))
	re = regexp.MustCompile(`([a-z])([A-Z])`)
	s = re.ReplaceAllString(s, strings.ToLower(`$1-$2`))
	re = regexp.MustCompile(`^-`)
	s = re.ReplaceAllString(strings.ToLower(s), "")
	return s
}

// Obfuscate s, typically used for hiding primary key values when passed along in the URL.
// A port of CFWheels obfuscateParam() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func Obfuscate(s string) string {
	// Check to make sure s doesn't start with a "0".
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

// StripLinks removes all HTML links from s, leaving just the link text.
// A port of CFWheels StripLinks() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
func StripLinks(s string) string {
	re := regexp.MustCompile(`<a.*?>(.*?)</a>`)
	s = re.ReplaceAllString(s, "$1")
	return s
}

// StripTags removes all HTML tags from s.
// A port of CFWheels stripTags() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm
func StripTags(s string) string {
	re := regexp.MustCompile(`<\ *[a-z].*?>`)
	s = re.ReplaceAllString(s, "")
	re = regexp.MustCompile(`<\ */\ *[a-z].*?>`)
	s = re.ReplaceAllString(s, "")
	return s
}

// TimeDistance describes the difference between two time.Time values.
// To include number of seconds in the description set sec to true.
// A port of CFWheels distanceOfTimeInWords(), timeAgoInWords() and timeUntilInWords() helpers.
// To simulate timeAgoInWords() make to time.Now().
// To simulate timeUntilInWords() make from time.Now().
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func TimeDistance(from time.Time, to time.Time, sec bool) string {
	delta := to.Sub(from)
	sd := int(delta.Seconds())
	md := int(delta.Minutes())
	hd := int(delta.Hours())
	var days, months, years int
	switch {
	case md <= 1:
		if sec != true {
			switch {
			case sd < 60:
				return "less than a minute"
			default:
				return "1 minute"
			}
		} else {
			switch {
			case sd < 5:
				return "less than 5 seconds"
			case sd < 10:
				return "less than 10 seconds"
			case sd < 20:
				return "less than 20 seconds"
			case sd < 40:
				return "half a minute"
			default:
				return "1 minute"
			}
		}
	case md < 45:
		return fmt.Sprintf("%d minutes", md)
	case md < 90:
		return "about 1 hour"
	case md < 1440:
		return fmt.Sprintf("about %d hours", hd)
	case md < 2880:
		return "1 day"
	case md < 43200:
		days = hd / 24
		return fmt.Sprintf("%d days", days)
	case md < 86400:
		return "about 1 month"
	case md < 525600:
		months = hd / 730
		return fmt.Sprintf("%d months", months)
	case md < 657000:
		return "about 1 year"
	case md < 919800:
		return "over 1 year"
	case md < 1051200:
		return "almost 2 years"
	default:
		years = md / 525600
		return fmt.Sprintf("over %d years", years)
	}
}

// Truncate s to the specified number and replaces the last characters with the specified new string (which defaults to "...").
// A port of CFWheels truncate() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func Truncate(s string, new string, n int) string {
	if new == "" {
		new = "..."
	}
	if len(s) <= n {
		return s
	}
	return s[0:n-len(new)] + new
}

// WordTruncate truncates s to the specified number of words and replaces the remaining characters with the specified
// new string (which defaults to "..."). A port of CFWheels wordTruncate() helper.
// A port of CFWheels wordTruncate() helper.
// CFML source: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm
func WordTruncate(s string, new string, n int) string {
	if new == "" {
		new = "..."
	}
	words := strings.Fields(s)
	// When there are fewer (or same) words in the string than the number to be truncated we can just return it unchanged.
	if len(words) >= utf8.RuneCountInString(s) {
		return s
	}
	var rv string
	for i, w := range words {
		if i+1 >= n {
			rv += w
			break
		}
		rv += w + " "
	}
	rv += new
	return rv
}

// reverseInt reverses an integer.
// An i of 12345 will return 54321.
func reverseInt(i int) (int, error) {
	// credit: Wade73, http://stackoverflow.com/questions/35972561/reverse-int-golang
	itoa := strconv.Itoa(i)
	var str string
	for x := len(itoa); x > 0; x-- {
		str += string(itoa[x-1])
	}
	rev, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return rev, nil
}
