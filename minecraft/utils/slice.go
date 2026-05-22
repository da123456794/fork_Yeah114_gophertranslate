package utils

import "fmt"

// ConvertSlice converts a slice of type From to a slice of type To using the provided converter function.
func ConvertSlice[From any, To any](
	from []From,
	converter func(from From) To,
) []To {
	to := make([]To, len(from))
	for i, v := range from {
		to[i] = converter(v)
	}
	return to
}

// ConvertSliceWithError converts a slice of type From to a slice of type To using the provided converter function. If the converter function returns an error for any element, the conversion is aborted and the error is returned.
func ConvertSliceWithError[From any, To any](
	from []From,
	converter func(from From) (To, error),
) ([]To, error) {
	to := make([]To, len(from))
	for i, v := range from {
		var err error
		to[i], err = converter(v)
		if err != nil {
			return nil, fmt.Errorf("ConvertSliceWithError: failed to convert element at index %d: %w", i, err)
		}
	}
	return to, nil
}
