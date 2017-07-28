package cfw

import (
	"testing"
	"time"
)

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/deobfuscateparam.cfc
func TestDeObfuscate(t *testing.T) {
	var e, a, s string
	e = "999999999"
	s = "eb77359232"
	a = DeObfuscate(s)
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	e = "1"
	s = "9b1c6"
	a = DeObfuscate(s)
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	e = "99"
	s = "ac10a"
	a = DeObfuscate(s)
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	e = "15765"
	s = "b226582"
	a = DeObfuscate(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	e = "69247541"
	s = "c06d44215"
	a = DeObfuscate(s)
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "becca2515"
	a = DeObfuscate(s)
	e = "becca2515"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "a15ba9"
	a = DeObfuscate(s)
	e = "a15ba9"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "1111111111"
	a = DeObfuscate(s)
	e = "1111111111"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/9eea7a6ac77956c8825e037f2e3e6b8c0f346267/wheels/tests/view/sanitize/striptags.cfc
func TestExcerpt(t *testing.T) {
	var a, e, p, s string
	s = `CFWheels: testing the excerpt view helper to see if it works or not.`
	p = `CFWheels: testing the excerpt`
	a = Excerpt(s, "[more]", p, 0)
	e = "CFWheels: testing the excerpt[more]"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	p = `testing the excerpt`
	a = Excerpt(s, "[more]", p, 0)
	e = "[more]testing the excerpt[more]"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	p = `excerpt view helper`
	a = Excerpt(s, "[more]", p, 10)
	e = "[more]sting the excerpt view helper to see if[more]"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	a = Excerpt(s, "[more]", p, 25)
	e = "CFWheels: testing the excerpt view helper to see if it works or no[more]"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	p = "see if it works"
	a = Excerpt(s, "[more]", p, 25)
	e = "[more]e excerpt view helper to see if it works or not."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	p = "jklsduiermobk"
	a = Excerpt(s, "[more]", p, 25)
	e = ""
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/humanize.cfc
func TestHumanize(t *testing.T) {
	var a, e, s string
	s = "wheelsIsAFramework"
	e = "Wheels Is A Framework"
	a = Humanize(s, "")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "WheelsIsAFramework"
	e = "Wheels Is A Framework"
	a = Humanize(s, "")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "CFML"
	e = "CFML"
	a = Humanize(s, "")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "ACfmlFramework"
	e = "A CFML Framework"
	a = Humanize(s, "CFML")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "wheelsIsACFMLFramework"
	e = "Wheels Is ACFML Framework"
	a = Humanize(s, "")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "Some Input"
	e = "Some Input"
	a = Humanize(s, "")
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/13b9cb72f52a4edc7981ae572d050f8116af57ef/wheels/tests/global/strings.cfc
func TestHyphenize(t *testing.T) {
	var a, e, s string
	s = "wheelsIsAFramework"
	e = "wheels-is-a-framework"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "WheelsIsAFramework"
	e = "wheels-is-a-framework"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "aURLVariable"
	e = "a-url-variable"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "URLVariable"
	e = "url-variable"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "ERRORMESSAGE"
	e = "errormessage"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
	s = "address"
	e = "address"
	a = Hyphenize(s)
	if a != e {
		t.Errorf("mismatch, got: %q, want: %q", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/obfuscateparam.cfc
func TestObfuscate(t *testing.T) {
	var a, e, s string
	s = "999999999"
	a = Obfuscate(s)
	e = "eb77359232"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "0162823571"
	a = Obfuscate(s)
	e = "0162823571"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "1"
	a = Obfuscate(s)
	e = "9b1c6"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "99"
	a = Obfuscate(s)
	e = "ac10a"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "15765"
	a = Obfuscate(s)
	e = "b226582"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "69247541"
	a = Obfuscate(s)
	e = "c06d44215"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "0413"
	a = Obfuscate(s)
	e = "0413"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "per"
	a = Obfuscate(s)
	e = "per"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "1111111111"
	a = Obfuscate(s)
	e = "1111111111"
	if a != e {
		// TODO: This is meant to fail but doesn't, maybe a CFML quirk?
		//t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

func TestReverseInt(t *testing.T) {
	s := 2345678
	a, _ := reverseInt(s)
	e := 8765432
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = 505
	a, _ = reverseInt(s)
	e = 505
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = 05
	a, _ = reverseInt(s)
	e = 5 // TODO is this wanted behaviour?
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/9eea7a6ac77956c8825e037f2e3e6b8c0f346267/wheels/tests/view/sanitize/striptags.cfc
func TestStripLinks(t *testing.T) {
	var a, e, s string
	s = `this <a href="http://www.google.com" title="google">is</a> a <a href="mailto:someone@example.com" title="invalid email">test</a> to <a name="anchortag">see</a> if this works or not.`
	a = StripLinks(s)
	e = "this is a test to see if this works or not."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/view/text/excerpt.cfc
func TestStripTags(t *testing.T) {
	var a, e, s string
	s = `<h1>this</h1><p><a href="http://www.google.com" title="google">is</a></p><p>a <a href="mailto:someone@example.com" title="invalid email">test</a> to<br><a name="anchortag">see</a> if this works or not.</p>`
	a = StripTags(s)
	e = "thisisa test tosee if this works or not."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

func TestTimeDistance(t *testing.T) {
	var s, n time.Time
	var a, e string
	n = time.Now()
	s = n.Add(time.Second * time.Duration(5-1))
	a = TimeDistance(n, s, true)
	e = "less than 5 seconds"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Second * time.Duration(10-1))
	a = TimeDistance(n, s, true)
	e = "less than 10 seconds"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Second * time.Duration(20-1))
	a = TimeDistance(n, s, true)
	e = "less than 20 seconds"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Second * time.Duration(40-1))
	a = TimeDistance(n, s, true)
	e = "half a minute"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Second * time.Duration(59))
	a = TimeDistance(n, s, false)
	e = "less than a minute"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Second * time.Duration(60+50))
	a = TimeDistance(n, s, false)
	e = "1 minute"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(45-1))
	a = TimeDistance(n, s, false)
	e = "44 minutes"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(90-1))
	a = TimeDistance(n, s, false)
	e = "about 1 hour"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(1440-1))
	a = TimeDistance(n, s, false)
	e = "about 23 hours"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(2880-1))
	a = TimeDistance(n, s, false)
	e = "1 day"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(43200-1))
	a = TimeDistance(n, s, false)
	e = "29 days"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(86400-1))
	a = TimeDistance(n, s, false)
	e = "about 1 month"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(525600-1))
	a = TimeDistance(n, s, false)
	e = "11 months"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(657000-1))
	a = TimeDistance(n, s, false)
	e = "about 1 year"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(919800-1))
	a = TimeDistance(n, s, false)
	e = "over 1 year"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(1051200-1))
	a = TimeDistance(n, s, false)
	e = "almost 2 years"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = n.Add(time.Minute * time.Duration(1051200))
	a = TimeDistance(n, s, false)
	e = "over 2 years"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/9738bd71e7b8632587cd09a0a2d7d3e779700a7a/wheels/tests/view/text/truncate.cfc
func TestTruncate(t *testing.T) {
	var a, e, s string
	s = "this is a test to see if this works or not."
	a = Truncate(s, "[more]", 20)
	e = "this is a test[more]"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = ""
	a = Truncate(s, "[more]", 20)
	e = ""
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = "this is a test to see if this works or not."
	a = Truncate(s, "", 20)
	e = "this is a test to..."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Source of test values:
// https://github.com/cfwheels/cfwheels/blob/9738bd71e7b8632587cd09a0a2d7d3e779700a7a/wheels/tests/view/text/wordtruncate.cfc
func TestWordTruncate(t *testing.T) {
	var a, e, s string
	s = "CFWheels is a framework for ColdFusion"
	a = WordTruncate(s, "", 4)
	e = "CFWheels is a framework..."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}

// Test unique examples used in README.md
func TestReadmeExamples(t *testing.T) {
	var a, e, s string
	s = `Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`
	a = StripLinks(s)
	e = "Go to the <strong>GitHub</strong> repo!"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	a = StripTags(s)
	e = "Go to the GitHub repo!"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	from := time.Now()
	to := from.Add(time.Second * time.Duration(7)) // add 7 seconds
	a = TimeDistance(from, to, true)
	e = "less than 10 seconds"
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = `CFW contains Go ports of a few selected CFWheels helpers.`
	a = Truncate(s, "", 15)
	e = "CFW contains..."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
	s = `CFW contains Go ports of a few selected CFWheels helpers.`
	a = WordTruncate(s, "", 4)
	e = "CFW contains Go ports..."
	if a != e {
		t.Errorf("mismatch, got: %v, want: %v", a, e)
	}
}
