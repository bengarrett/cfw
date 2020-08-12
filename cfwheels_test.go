package cfw

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func ExampleDeObfuscate() {
	fmt.Println(DeObfuscate("9b1c6"))
	// Output: 1
}

func TestDeObfuscate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		// values:
		// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/deobfuscateparam.cfc
		{"empty", "", ""},
		{"ok 1", "9b1c6", "1"},
		{"ok 2", "eb77359232", "999999999"},
		{"ok 3", "ac10a", "99"},
		{"ok 4", "b226582", "15765"},
		{"ok 5", "c06d44215", "69247541"},
		{"ok 6", "ac10a", "99"},
		{"error 1", "becca2515", "becca2515"},
		{"error 2", "a15ba9", "a15ba9"},
		{"error 3", "1111111111", "1111111111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeObfuscate(tt.s); got != tt.want {
				t.Errorf("DeObfuscate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleExcerpt() {
	pangram := "The quick brown fox jumps over the lazy dog"
	fmt.Println(Excerpt(pangram, "...", "The quick brown fox", 0))
	fmt.Println(Excerpt(pangram, "...", "", 19))
	fmt.Println(Excerpt(pangram, "🦊", "The quick brown ", 0))
	// Output: The quick brown fox...
	// The quick brown fox...
	// The quick brown 🦊
}

func TestExcerpt(t *testing.T) {
	type args struct {
		s       string
		replace string
		phrase  string
		n       int
	}
	const s = "CFWheels: testing the excerpt view helper to see if it works or not."
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{}, "..."},
		{"1", args{s, "[more]", "CFWheels: testing the excerpt", 0}, "CFWheels: testing the excerpt[more]"},
		{"2", args{s, "[more]", "testing the excerpt", 0}, "[more]testing the excerpt[more]"},
		{"3", args{s, "[more]", "excerpt view helper", 10}, "[more]sting the excerpt view helper to see if[more]"},
		{"4", args{s, "[more]", "excerpt view helper", 25}, "CFWheels: testing the excerpt view helper to see if it works or no[more]"},
		{"5", args{s, "[more]", "see if it works", 25}, "[more]e excerpt view helper to see if it works or not."},
		{"6", args{s, "[more]", "jklsduiermobk", 25}, ""},
		{"utf8", args{"The quick brown 🦊 jumps over the lazy 🐕", "💬", "brown 🦊", 0}, "💬brown 🦊💬"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Excerpt(tt.args.s, tt.args.replace, tt.args.phrase, tt.args.n); got != tt.want {
				t.Errorf("Excerpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHumanize() {
	fmt.Println(Humanize("helloWorldExample", []string{}...))
	fmt.Println(Humanize("golangModOrVgo?", []string{"MOD", "VGO"}...))
	// Output: Hello World Example
	// Golang MOD Or VGO?
}

func TestHumanize(t *testing.T) {
	type args struct {
		s      string
		except []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// values https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/humanize.cfc
		{"empty", args{"", nil}, ""},
		{"lcase", args{"wheelsIsAFramework", nil}, "Wheels Is A Framework"},
		{"title", args{"WheelsIsAFramework", nil}, "Wheels Is A Framework"},
		{"ucase", args{"CFML", []string{"CFML"}}, "CFML"},
		{"except", args{"ACfmlFramework", []string{"CFML"}}, "A CFML Framework"},
		{"err 1", args{"wheelsIsACFMLFramework", nil}, "Wheels Is ACFML Framework"},
		{"same", args{"Some Input", nil}, "Some Input"},
		{"emoji", args{"theQuickBrown🦊JumpsOverTheLazy🐕", nil}, "The Quick Brown🦊 Jumps Over The Lazy🐕"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Humanize(tt.args.s, tt.args.except...); got != tt.want {
				t.Errorf("Humanize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHyphenize() {
	fmt.Println(Hyphenize("aTourOfGo"))
	// Output: a-tour-of-go
}

func TestHyphenize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		// values https://github.com/cfwheels/cfwheels/blob/13b9cb72f52a4edc7981ae572d050f8116af57ef/wheels/tests/global/strings.cfc
		{"empty", "", ""},
		{"ok 1", "wheelsIsAFramework", "wheels-is-a-framework"},
		{"ok 2", "WheelsIsAFramework", "wheels-is-a-framework"},
		{"ok 3", "aURLVariable", "a-url-variable"},
		{"ok 4", "URLVariable", "url-variable"},
		{"ucase", "ERRORMESSAGE", "errormessage"},
		{"lcase", "address", "address"},
		{"emoji", "TheQuickBrown🦊JumpsOverTheLazy🐕", "the-quick-brown🦊-jumps-over-the-lazy🐕"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hyphenize(tt.s); got != tt.want {
				t.Errorf("Hyphenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleObfuscate() {
	fmt.Println(Obfuscate("1"))
	fmt.Println(Obfuscate("5551234"))
	// Output: 9b1c6
	// b3da865e
}

func TestObfuscate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		// values https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/global/public/obfuscateparam.cfc
		{"empty", "", ""},
		{"", "999999999", "eb77359232"},
		{"", "0162823571", "0162823571"},
		{"", "1", "9b1c6"},
		{"", "99", "ac10a"},
		{"", "15765", "b226582"},
		{"", "69247541", "c06d44215"},
		{"", "0413", "0413"},
		{"", "per", "per"},
		// in CFWheels this test fails but in Go it returns a429646180a
		// {"", "1111111111", "1111111111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Obfuscate(tt.s); got != tt.want {
				t.Errorf("Obfuscate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseInt(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		want    int
		wantErr bool
	}{
		{"zero", 0, 0, false},
		{"ok", 2345678, 8765432, false},
		{"505", 505, 505, false},
		{"leading 0", 0005, 5, false}, // expected behaviour?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverseInt(tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("reverseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("reverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleStripLinks() {
	fmt.Println(StripLinks(`<a href="https://golang.org">The Go Programming Language</a>.`))
	// Output: The Go Programming Language.
}

func TestStripLinks(t *testing.T) {
	// values: https://github.com/cfwheels/cfwheels/blob/9eea7a6ac77956c8825e037f2e3e6b8c0f346267/wheels/tests/view/sanitize/striptags.cfc
	str := `this <a href="http://www.google.com" title="google">is</a> a <a href="mailto:someone@example.com" title="invalid email">` +
		`test</a> to <a name="anchortag">see</a> if this works or not.`
	emoji := `The quick <b><a href="https://example.com">brown 🦊</a></b> jumps over the lazy 🐕`
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"string", str, "this is a test to see if this works or not."},
		{"emoji", emoji, "The quick <b>brown 🦊</b> jumps over the lazy 🐕"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripLinks(tt.s); got != tt.want {
				t.Errorf("StripLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleStripTags() {
	fmt.Println(StripTags(`<h1><a href="https://golang.org">The Go Programming Language</a>.</h1>`))
	// Output: The Go Programming Language.
}

func TestStripTags(t *testing.T) {
	// values:
	// https://github.com/cfwheels/cfwheels/blob/1c3b9d6db79cdfbfbe49ae6816f6dc96262ccf82/wheels/tests/view/text/excerpt.cfc
	str := `<h1>this</h1><p><a href="http://www.google.com" title="google">is</a></p><p>a ` +
		`<a href="mailto:someone@example.com" title="invalid email">test</a> to<br><a name="anchortag">see</a> if this works or not.</p>`
	emoji := `The quick <b><a href="https://example.com">brown 🦊</a></b> jumps over the lazy 🐕`
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"string", str, "thisisa test tosee if this works or not."},
		{"emoji", emoji, "The quick brown 🦊 jumps over the lazy 🐕"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripTags(tt.s); got != tt.want {
				t.Errorf("StripTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTimeDistance() {
	const layout = "2006 Jan 2"
	f, err := time.Parse(layout, "2000 Jan 1")
	if err != nil {
		log.Fatal(err)
	}
	t, err := time.Parse(layout, "2020 Jun 30")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(TimeDistance(f, t, false))
	// Output: over 20 years
}

func ExampleTimeDistance_seconds() {
	f, err := time.Parse(time.RFC3339, "2006-01-02T15:04:06+07:00")
	if err != nil {
		log.Fatal(err)
	}
	t, err := time.Parse(time.RFC3339, "2006-01-02T15:04:10+07:00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(TimeDistance(f, t, true))
	// Output: less than 5 seconds
}

func TestTimeDistance(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
		sec  bool
	}
	var n = time.Now()
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{}, "less than a minute"},
		{"<1m", args{n, n.Add(time.Second * time.Duration(5-1)), false}, "less than a minute"},
		{"<10s", args{n, n.Add(time.Second * time.Duration(10-1)), true}, "less than 10 seconds"},
		{"<20s", args{n, n.Add(time.Second * time.Duration(20-1)), true}, "less than 20 seconds"},
		{"1/2min", args{n, n.Add(time.Second * time.Duration(40-1)), true}, "half a minute"},
		{"<1min", args{n, n.Add(time.Second * time.Duration(59)), false}, "less than a minute"},
		{"1min", args{n, n.Add(time.Second * time.Duration(60+50)), false}, "1 minute"},
		{"44min", args{n, n.Add(time.Minute * time.Duration(45-1)), false}, "44 minutes"},
		{"1h", args{n, n.Add(time.Minute * time.Duration(90-1)), false}, "about 1 hour"},
		{"23h", args{n, n.Add(time.Minute * time.Duration(1440-1)), false}, "about 23 hours"},
		{"1d", args{n, n.Add(time.Minute * time.Duration(2880-1)), false}, "1 day"},
		{"29d", args{n, n.Add(time.Minute * time.Duration(43200-1)), false}, "29 days"},
		{"1m", args{n, n.Add(time.Minute * time.Duration(86400-1)), false}, "about 1 month"},
		{"11m", args{n, n.Add(time.Minute * time.Duration(525600-1)), false}, "11 months"},
		{"1y", args{n, n.Add(time.Minute * time.Duration(657000-1)), false}, "about 1 year"},
		{">1y", args{n, n.Add(time.Minute * time.Duration(919800-1)), false}, "over 1 year"},
		{"2y", args{n, n.Add(time.Minute * time.Duration(1051200-1)), false}, "almost 2 years"},
		{">2y", args{n, n.Add(time.Minute * time.Duration(1051200)), false}, "over 2 years"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeDistance(tt.args.from, tt.args.to, tt.args.sec); got != tt.want {
				t.Errorf("TimeDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTruncate() {
	fmt.Println(Truncate("Go is an open source programming language", "", 8))
	fmt.Println(Truncate("Go is an open source programming language", "?", 6))
	// Output: Go is...
	// Go is?
}

func TestTruncate(t *testing.T) {
	type args struct {
		s       string
		replace string
		n       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// values: https://github.com/cfwheels/cfwheels/blob/9738bd71e7b8632587cd09a0a2d7d3e779700a7a/wheels/tests/view/text/truncate.cfc
		{"empty", args{}, ""},
		{"ok1", args{"this is a test to see if this works or not.", "[more]", 20}, "this is a test[more]"},
		{"err1", args{"", "[more]", 20}, ""},
		{"ok2", args{"this is a test to see if this works or not.", "", 20}, "this is a test to..."},
		{"emoji", args{"The quick brown 🦊 jumps over the lazy 🐕", "💬", 21}, "The quick brown 🦊💬"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.s, tt.args.replace, tt.args.n); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleWordTruncate() {
	fmt.Println(WordTruncate("Go is an open source programming language", "", 2))
	fmt.Println(WordTruncate("Go is an open source programming language", "?", 2))
	// Output: Go is...
	// Go is?
}

func TestWordTruncate(t *testing.T) {
	type args struct {
		s       string
		replace string
		n       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// values: https://github.com/cfwheels/cfwheels/blob/9738bd71e7b8632587cd09a0a2d7d3e779700a7a/wheels/tests/view/text/wordtruncate.cfc
		{"empty", args{}, ""},
		{"ok", args{"CFWheels is a framework for ColdFusion", "", 4}, "CFWheels is a framework..."},
		{"emoji", args{"The quick brown 🦊 jumps over the lazy 🐕", "💬", 4}, "The quick brown 🦊💬"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordTruncate(tt.args.s, tt.args.replace, tt.args.n); got != tt.want {
				t.Errorf("WordTruncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test unique examples used in README.md.
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
