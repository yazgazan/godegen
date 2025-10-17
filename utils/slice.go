package utils

type Slice struct{}

func (Slice) New(values ...any) []any {
	return values
}

func (Slice) Append(v any, slice []any) []any {
	return append(slice, v)
}
