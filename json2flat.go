package json2flat

import (
	"iter"
	"strconv"
)

type Value interface {
	AsAny() any
}

type MapToFlatMap func(map[string]any) map[string]Value
type MapToFlatAny func(map[string]any) map[string]any

func (m MapToFlatMap) ToMapToFlatAny() MapToFlatAny {
	buf := map[string]any{}
	return func(original map[string]any) map[string]any {
		clear(buf)

		var converted map[string]Value = m(original)
		for key, val := range converted {
			var a any = val.AsAny()
			buf[key] = a
		}
		return buf
	}
}

func KeyValToPairs(
	delim string,
	key string,
	val any,
) iter.Seq2[string, Value] {
	switch t := val.(type) {

	case bool:
		return func(yield func(string, Value) bool) {
			yield(key, Boolean(t))
		}

	case float64:
		return func(yield func(string, Value) bool) {
			yield(key, Double(t))
		}

	case string:
		return func(yield func(string, Value) bool) {
			yield(key, String(t))
		}

	case nil:
		return func(yield func(string, Value) bool) {
			yield(key, Null(struct{}{}))
		}

	case map[string]any:
		return func(yield func(string, Value) bool) {
			for subKey, val := range t {
				var pairs iter.Seq2[string, Value] = KeyValToPairs(
					delim,
					key+delim+subKey,
					val,
				)

				for s, v := range pairs {
					if !yield(s, v) {
						return
					}
				}
			}
		}

	case []any:
		return func(yield func(string, Value) bool) {
			for ix, val := range t {
				var pairs iter.Seq2[string, Value] = KeyValToPairs(
					delim,
					key+delim+strconv.Itoa(ix),
					val,
				)

				for s, v := range pairs {
					if !yield(s, v) {
						return
					}
				}
			}
		}

	default:
		panic(t)

	}
}

func MapToFlatMapNew(delim string) MapToFlatMap {
	buf := map[string]Value{}
	return func(original map[string]any) map[string]Value {
		clear(buf)

		for key, val := range original {
			var pairs iter.Seq2[string, Value] = KeyValToPairs(
				delim,
				key,
				val,
			)
			for k, v := range pairs {
				buf[k] = v
			}
		}
		return buf
	}
}

const DelimDefault string = "_"

func MapToFlatMapDefault() MapToFlatMap { return MapToFlatMapNew(DelimDefault) }

type Boolean bool

func (b Boolean) AsAny() any {
	var raw bool = bool(b)
	return raw
}
func (b Boolean) AsValue() Value { return b }

type Double float64

func (d Double) AsAny() any {
	var raw float64 = float64(d)
	return raw
}
func (d Double) AsValue() Value { return d }

type String string

func (s String) AsAny() any {
	var raw string = string(s)
	return raw
}
func (s String) AsValue() Value { return s }

type Null struct{}

func (n Null) AsAny() any     { return nil }
func (n Null) AsValue() Value { return n }
