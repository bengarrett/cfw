# CFW

## Package cfw contains ports of a few selected CFWheels helpers that are used for string manipulation and have no Go equivalent

[![Build Status](https://travis-ci.org/bengarrett/cfw.svg?branch=master)](https://travis-ci.org/bengarrett/cfw) [![Go Report Card](https://goreportcard.com/badge/github.com/bengarrett/cfw)](https://goreportcard.com/report/github.com/bengarrett/cfw) [![GoDoc](https://godoc.org/github.com/bengarrett/cfw?status.svg)](https://godoc.org/github.com/bengarrett/cfw)

_[CFWheels](https://cfwheels.org/) is an open source [CFML (ColdFusion Markup Language)](http://lucee.org/) framework inspired by Ruby on Rails that provides fast application development, a great organization system for your code, and is just plain fun to use._

I created this small package as I have been migrating old CFWheels web applications over to Go. They had a few CFWheels dependencies that needed to be recreated to ensure a smooth transition.

## Import

```Go
import "github.com/bengarrett/cfw"
```

### `DeObfuscate`

```Go
println(cfw.DeObfuscate("b226582"))
```

Would return `15765`

### `Excerpt`

```Go
s := "CFWheels: testing the excerpt view helper to see if it works or not."
p := "excerpt view helper"
println(cfw.Excerpt(s, "[more]", p, 10))`
```

Would return `[more]sting the excerpt view helper to see if[more]`

### `Humanize`

```Go
println(cfw.Humanize("wheelsIsAFramework,""))
```

Would return `Wheels Is A Framework`

### `Hyphenize`

```Go
println(cfw.Hyphenize("wheelsIsAFramework"))
```

Would return `wheels-is-a-framework`

### `Obfuscate`

```Go
println(cfw.Obfuscate("15765"))
```

Would return `b226582`

### `StripLinks`

```Go
l := `Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`
println(cfw.StripLinks(l))
```

Would return `Go to the <strong>GitHub</strong> repo!`

### `StripTags`

```Go
l := `Go to the <strong><a href="https://github.com/bengarrett/cfw">GitHub</a></strong> repo!`
println(cfw.StripTags(l))
```

Would return `Go to the GitHub repo!`

### `TimeDistance`

```Go
from := time.Now()
to := from.Add(time.Second * time.Duration(7)) // add 7 seconds
println(cfw.TimeDistance(from, to, true))
```

Would return `less than 10 seconds`

### `Truncate`

```Go
s := `CFW contains Go ports of a few selected CFWheels helpers.`
println(cfw.Truncate(s, "", 15))

```

Would return `CFW contains...`

### `WordTruncate`

```Go
s := `CFW contains Go ports of a few selected CFWheels helpers.`
println(cfw.WordTruncate(s, "", 4))

```

Would return `CFW contains Go ports...`