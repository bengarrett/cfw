package main

import (
	"fmt"
	"time"

	"github.com/bengarrett/cfw"
)

func main() {
	// Obfuscate / DeObfuscate
	fmt.Println(cfw.Obfuscate("1234"))
	fmt.Println(cfw.DeObfuscate("a4363c"))

	// Excerpt
	chars := 10
	fmt.Println(cfw.Excerpt(
		"CFWheels: testing the excerpt view helper to see if it works or not.",
		"[more]", "excerpt view helper", chars))

	// Humanize
	fmt.Println(cfw.Humanize("wheelsIsAFramework", ""))

	// Hyphenize
	fmt.Println(cfw.Hyphenize("wheelsIsAFramework"))

	// StripLinks
	fmt.Println(cfw.StripLinks(
		`Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

	// StripTags
	fmt.Println(cfw.StripTags(
		`Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

	// TimeDistance
	const seven = time.Second * time.Duration(7)

	fmt.Println(cfw.TimeDistance(time.Now(), time.Now().Add(seven), true))

	// Truncate
	chars = 15
	fmt.Println(cfw.Truncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", "", chars))

	// WordTruncate
	words := 4
	fmt.Println(cfw.WordTruncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", "", words))

	words = 3
	fmt.Println(cfw.WordTruncate(
		"CFW contains Go ports of a few selected CFWheels helpers.", " 🥰", words))
}
