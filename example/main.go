package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bengarrett/cfw"
)

func main() {
	w := os.Stdout

	// Obfuscate / DeObfuscate
	fmt.Fprintln(w, cfw.Obfuscate("1234"))
	fmt.Fprintln(w, cfw.DeObfuscate("a4363c"))

	// Excerpt
	chars := 10
	fmt.Fprintln(w, cfw.Excerpt(
		"CFWheels: testing the excerpt view helper to see if it works or not.",
		"[more]", "excerpt view helper", chars))

	// Humanize
	fmt.Fprintln(w, cfw.Humanize("wheelsIsAFramework", ""))

	// Hyphenize
	fmt.Fprintln(w, cfw.Hyphenize("wheelsIsAFramework"))

	// StripLinks
	fmt.Fprintln(w, cfw.StripLinks(
		`Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

	// StripTags
	fmt.Fprintln(w, cfw.StripTags(
		`Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

	// TimeDistance
	const seven = time.Second * time.Duration(7)

	fmt.Fprintln(w, cfw.TimeDistance(time.Now(), time.Now().Add(seven), true))

	// Truncate
	chars = 15
	fmt.Fprintln(w, cfw.Truncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", "", chars))

	// WordTruncate
	words := 4
	fmt.Fprintln(w, cfw.WordTruncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", "", words))

	words = 3
	fmt.Fprintln(w, cfw.WordTruncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", " 🥰", words))
}
