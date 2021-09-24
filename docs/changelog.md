## v1.3
- Go v1.17 usage.
- New `ReverseInt()` function.
- Linter updates and fixes.
- GitHub repository updates and fixes.

## v1.2

- For use clarity, changed the type of `Humanize()` `except` argument from `[]string` to `...string`.<br>
  All uses now require that you pass on `[]string{}...` values to the argument.
- `reverseInt()` uses Go v1.13 style error wrap.
- Removed magic numbers and instead use named constants for `int` values.
- Fixed shadow declarations.
- Fixed ExampleTimeDistance_seconds() test that ignored a timezone error.

## v1.1

- `Humanize()` argument type change: `except []string`.
- Fixed some tests using incorrect argument names.
