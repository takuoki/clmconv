# clmconv

[![CircleCI](https://circleci.com/gh/takuoki/clmconv.svg?style=shield&circle-token=0d4dc479afa80d4399bb7a457584142d37d3f50e)](https://circleci.com/gh/takuoki/clmconv)
[![codecov](https://codecov.io/gh/takuoki/clmconv/branch/master/graph/badge.svg)](https://codecov.io/gh/takuoki/clmconv)
[![GoDoc](https://godoc.org/github.com/takuoki/clmconv?status.svg)](https://godoc.org/github.com/takuoki/clmconv)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang package for converting to spreadsheet column alphabet or integer.

## Usage

### Default converter

```go
i, err := clmconv.Atoi("A") // i = 0
```

```go
i := clmconv.MustAtoi("a") // i = 0
```

```go
a := clmconv.Itoa(0) // a = "A"
```

### Custom converter

```go
converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
```

```go
i, err := converter.Atoi("A") // i = 1
```

```go
i := converter.MustAtoi("a") // i = 1
```

```go
a := converter.Itoa(1) // a = "a"
```

## Example

| Alphabet | Integer | Integer (WithStartFromOne) |
| :------- | ------: | -------------------------: |
| A        |       0 |                          1 |
| B        |       1 |                          2 |
| Z        |      25 |                         26 |
| AA       |      26 |                         27 |
| ZZ       |     701 |                        702 |
| ABC      |     730 |                        731 |
| ABCDE    |  494264 |                     494265 |

## References

- [Blog: Library to convert Spreadsheet column string to integer](https://medium.com/veltra-engineering/library-to-convert-spreadsheet-column-string-to-integer-72a6562c5f5c) (Japanese)
