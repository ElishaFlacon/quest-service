package utils

func MapToPrimitiveArray[I, O any](ts []*I, fn func(*I) O) []O {
	result := make([]O, len(ts))
	for i := range ts {
		result[i] = fn(ts[i])
	}
	return result
}
