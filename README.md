# clmconv

[![CircleCI](https://circleci.com/gh/takuoki/clmconv.svg?style=svg&circle-token=0d4dc479afa80d4399bb7a457584142d37d3f50e)](https://circleci.com/gh/takuoki/clmconv)

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
