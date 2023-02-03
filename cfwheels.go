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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	ellipsis     = "..."
	hexadecimal  = 16
	obfuscateXOR = 461
	obfuscateSum = 154
)

// Deobfuscate the obfuscated string, or return the original string.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// See: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L508
func DeObfuscate(s string) string {
	const checksum, decimal = 2, 10
	if len(s) < checksum {
		return s
	}
	if i, _ := strconv.Atoi(s); i > 0 {
		return s
	}
	// deobfuscate string
	num, err := strconv.ParseInt(s[checksum:], hexadecimal, 0)
	if err != nil {
		return s
	}
	num ^= obfuscateXOR
	baseNum := strconv.Itoa(int(num))
	l := len(baseNum) - 1
	value := ""
	for i := 0; i < l; i++ {
		f := baseNum[l-i:][:1]
		value += f
	}
	// create checks
	l = len(value)
	chksumTest := 0
	for i := 0; i < l; i++ {
		chr := value[i : i+1]
		n, err1 := strconv.Atoi(chr)
		if err1 != nil {
			return s
		}
		chksumTest += n
	}
	// run checks
	chksum, err := strconv.ParseInt(s[:2], hexadecimal, 0)
	if err != nil {
		return s
	}
	chksumX := strconv.FormatInt(chksum, decimal)
	chksumY := strconv.FormatInt(int64(chksumTest+obfuscateSum), decimal)
	if err := chksumX != chksumY; err {
		return s
	}

	return value
}

// Excerpt replaces n characters from s, which match the first instance of a given phrase.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// See: https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L68
func Excerpt(s, replace, phrase string, n int) string {
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
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/632ea90547da368cddd77cefe17f42a7eda871e0/wheels/global/util.cfm#L53
func Humanize(s string, except ...string) string {
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
	c := cases.Title(language.English, cases.NoLower)

	return c.String(s)
}

// Hyphenize converts camelCase strings to a lowercase hyphened string.
func Hyphenize(s string) string {
	s = regexp.MustCompile(`([A-Z][a-z])`).ReplaceAllString(s, strings.ToLower(`-$1`))
	s = regexp.MustCompile(`([a-z])([A-Z])`).ReplaceAllString(s, strings.ToLower(`$1-$2`))
	s = regexp.MustCompile(`^-`).ReplaceAllString(strings.ToLower(s), "")

	return s
}

// Obfuscate a numeric string to insecurely hide database primary key values when passed along a URL.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L483
func Obfuscate(s string) string {
	i, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	// confirm the first digit of i isn't a zero
	if s[0] == '0' {
		return s
	}
	reverse, err := ReverseInt(i)
	if err != nil {
		return s
	}
	l := len(s)
	a := int(math.Pow10(l) + float64(reverse))
	b := 0
	for i := 1; i <= l; i++ {
		// slice and sum the individual digits
		digit, err := strconv.Atoi(string(s[l-i]))
		if err != nil {
			return s
		}
		b += digit
	}
	// base64 conversion
	a ^= obfuscateXOR
	b += obfuscateSum

	return fmt.Sprintf("%s%s",
		strconv.FormatInt(int64(b), hexadecimal),
		strconv.FormatInt(int64(a), hexadecimal),
	)
}

// StripLinks removes all HTML links from a string leaving just the link text.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm#L3
func StripLinks(s string) string {
	return regexp.MustCompile(`<a.*?>(.*?)</a>`).ReplaceAllString(s, "$1")
}

// StripTags removes all HTML tags from a string.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/daa7c43fc993cab00f52cf8ac881e6cc93c02fe1/wheels/view/sanitize.cfm#L21
func StripTags(s string) string {
	x := regexp.MustCompile(`<\ *[a-z].*?>`).ReplaceAllString(s, "")

	return regexp.MustCompile(`<\ */\ *[a-z].*?>`).ReplaceAllString(x, "")
}

// TimeDistance describes the difference between two time values.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L112
func TimeDistance(from, to time.Time, seconds bool) string {
	delta := to.Sub(from)
	secs, mins, hrs := int(delta.Seconds()), int(delta.Minutes()), int(delta.Hours())

	const hours, days, months, year, years, twoyears = 1440, 43200, 525600, 657000, 919800, 1051200

	switch {
	case mins <= 1:
		if !seconds {
			return lessMin(secs)
		}

		return lessMinAsSec(secs)
	case mins < hours:
		return lessHours(mins, hrs)
	case mins < days:
		return lessDays(mins, hrs)
	case mins < months:
		return lessMonths(mins, hrs)
	case mins < year:
		return "about 1 year"
	case mins < years:
		return "over 1 year"
	case mins < twoyears:
		return "almost 2 years"
	default:
		y := mins / months

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

func lessHours(mins, hrs int) string {
	const parthour, abouthour, hours = 45, 90, 1440

	switch {
	case mins < parthour:
		return fmt.Sprintf("%d minutes", mins)
	case mins < abouthour:
		return "about 1 hour"
	case mins < hours:
		return fmt.Sprintf("about %d hours", hrs)
	default:
		return ""
	}
}

func lessDays(mins, hrs int) string {
	const day, days = 2880, 43200

	switch {
	case mins < day:
		return "1 day"
	case mins < days:
		const hoursinaday = 24
		d := hrs / hoursinaday

		return fmt.Sprintf("%d days", d)
	default:
		return ""
	}
}

func lessMonths(mins, hrs int) string {
	const month, months = 86400, 525600

	switch {
	case mins < month:
		return "about 1 month"
	case mins < months:
		const hoursinamonth = 730
		m := hrs / hoursinamonth

		return fmt.Sprintf("%d months", m)
	default:
		return ""
	}
}

// Truncate a string to the specified number and replace the trailing characters.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L20
func Truncate(s, replace string, n int) string {
	if replace == "" {
		replace = ellipsis
	}

	if utf8.RuneCountInString(s) <= n {
		return s
	}

	return s[0:n-utf8.RuneCountInString(replace)] + replace
}

// WordTruncate truncates a string to the specified number of words and replaces the trailing characters.
// This function is a port of a CFWheels framework function programmed in ColdFusion (CFML).
// https://github.com/cfwheels/cfwheels/blob/cf8e6da4b9a216b642862e7205345dd5fca34b54/wheels/global/misc.cfm#L40
func WordTruncate(s, replace string, n int) string {
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

// ReverseInt reverses an integer.
func ReverseInt(i int) (int, error) {
	// credit: Wade73
	// http://stackoverflow.com/questions/35972561/reverse-int-golang
	itoa, str := strconv.Itoa(i), ""
	for x := len(itoa); x > 0; x-- {
		str += string(itoa[x-1])
	}

	reverse, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("reverseInt %d: %w", i, err)
	}

	return reverse, nil
}
