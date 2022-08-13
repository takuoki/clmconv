// Package clmconv converts spreadsheet column alphabet to integer and vice versa.
package clmconv

const (
	runeOffsetUpper = 64
	runeOffsetLower = 96
)

// Converter represents a conversion object.
// If you need a conversion other than the default,
// you need create a new conversion object.
type Converter struct {
	offset     int
	runeOffset int
}

var defaultConverter = newDefault()

func newDefault() *Converter {
	return &Converter{
		offset:     1,
		runeOffset: runeOffsetUpper,
	}
}

// New creates a new Converter.
func New(opts ...Option) *Converter {
	c := newDefault()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

// Option is an option that configures Converter.
type Option interface {
	apply(*Converter)
}

type funcOption struct {
	f func(*Converter)
}

func (f *funcOption) apply(o *Converter) {
	f.f(o)
}

func newFuncOption(f func(*Converter)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithStartFromOne is an option to start the index from one.
// This option works for both `Atoi` and `Itoa`.
func WithStartFromOne() Option {
	return newFuncOption(func(c *Converter) {
		c.offset = 0
	})
}

// WithLowercase is an option to use lowercase letters as the output alphabet.
// This option works only with `Itoa`.
func WithLowercase() Option {
	return newFuncOption(func(c *Converter) {
		c.runeOffset = runeOffsetLower
	})
}
