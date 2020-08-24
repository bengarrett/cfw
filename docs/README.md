# cfw

![GitHub](https://img.shields.io/github/license/bengarrett/cfw?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/bengarrett/cfw)](https://goreportcard.com/report/github.com/bengarrett/cfw) 
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bengarrett/cfw)](https://github.com/bengarrett/cfw/releases)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/bengarrett/cfw)](https://pkg.go.dev/github.com/bengarrett/cfw)

#### Package cfw contains ports of a few selected CFWheels v1 helpers that are used for string manipulation and have no Go equivalent.

>[CFWheels](https://cfwheels.org/) is an open-source [CFML (ColdFusion Markup Language)](http://lucee.org/) framework inspired by Ruby on Rails that provides fast application development, a great organization system for your code, and is just plain fun to use.

I created this package as I had been migrating CFWheels web applications over to Go. They had a few dependencies that need recreation to ensure a smooth transition.

☕ [Find cfw useful, why not buy me a coffee?](https://www.buymeacoffee.com/4rtEGvUIY)

## Documentation

https://pkg.go.dev/github.com/bengarrett/cfw

## Usages

```Go
package main

import "github.com/bengarrett/cfw"

func main() {

    // Obfuscate / DeObfuscate
    println(cfw.Obfuscate("1234"))
    println(cfw.DeObfuscate("a4363c"))
    
    // Excerpt
    println(cfw.Excerpt(
	"CFWheels: testing the excerpt view helper to see if it works or not.",
        "[more]", "excerpt view helper", 10))

    // Humanize
    println(cfw.Humanize("wheelsIsAFramework", nil))

    // Hyphenize
    println(cfw.Hyphenize("wheelsIsAFramework"))

    // StripLinks
    println(cfw.StripLinks(
        `Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

    // StripTags
    println(cfw.StripTags(
        `Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`))

    // TimeDistance
    add7Seconds := from.Add(time.Second * time.Duration(7)
    println(cfw.TimeDistance(time.Now(), add7Seconds), true))

    // Truncate
    println(cfw.Truncate(
        "CFW contains Go ports of a few selected CFWheels helpers.", "", 15))

    // WordTruncate
    println(cfw.WordTruncate(
        "CFW contains Go ports of a few selected CFWheels helpers.", "", 4))
    println(cfw.WordTruncate(
        "CFW contains Go ports of a few selected CFWheels helpers.", " 🥰", 3))
}
```

Will print

```bash
# Obfuscate / DeObfuscate
a4363c
1234

# Excerpt
[more]sting the excerpt view helper to see if[more]

# Humanize
Wheels Is A Framework

# Hyphenize
wheels-is-a-framework

# StripLinks
Go to the <strong>GitHub</strong> repo!

# StripTags
Go to the GitHub repo!

# TimeDistance
less than 10 seconds

# Truncate
CFW contains...

# WordTruncate
CFW contains Go ports...
CFW contains Go 🥰
```

##### Copyright © 2020 [Ben Garrett](mailto:code.by.ben@gmail.com), [MIT License](https://pkg.go.dev/github.com/fluhus/godoc-tricks?tab=licenses)
