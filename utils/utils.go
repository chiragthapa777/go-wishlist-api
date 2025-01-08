package utils

// Map applies a transformation function to each element in the input slice and returns a new slice with transformed values.
func Map[T any, R any](input []T, transform func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce combines elements of the input slice into a single value using an accumulator function.
func Reduce[T any, R any](input []T, initial R, accumulator func(R, T) R) R {
	result := initial
	for _, v := range input {
		result = accumulator(result, v)
	}
	return result
}
