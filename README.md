# clmconv

[![CircleCI](https://circleci.com/gh/takuoki/clmconv.svg?style=shield&circle-token=0d4dc479afa80d4399bb7a457584142d37d3f50e)](https://circleci.com/gh/takuoki/clmconv)
[![codecov](https://codecov.io/gh/takuoki/clmconv/branch/master/graph/badge.svg)](https://codecov.io/gh/takuoki/clmconv)
[![GoDoc](https://godoc.org/github.com/takuoki/clmconv?status.svg)](https://godoc.org/github.com/takuoki/clmconv)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang package for converting to spreadsheet column alphabet or integer.

## Usage

```go
i, err := clmconv.Atoi("A") // i = 0
```

```go
i := clmconv.MustAtoi("A") // i = 0
```

```go
a := clmconv.Itoa(0) // a = "A"
```

## Example

| Alphabet | Integer |
|:---------|--------:|
| A        |       0 |
| B        |       1 |
| Z        |      25 |
| AA       |      26 |
| ZZ       |     701 |
| ABC      |     730 |
| ABCDE    |  494264 |

## References

- [Blog: Library to convert Spreadsheet column string to integer](https://medium.com/veltra-engineering/library-to-convert-spreadsheet-column-string-to-integer-72a6562c5f5c) (Japanese)
