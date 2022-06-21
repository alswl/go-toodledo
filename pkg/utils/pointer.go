package utils

func WrapListPointer[T any](list []T) []*T {
	var result []*T
	for i := range list {
		result = append(result, &list[i])
	}
	return result
}

func UnwrapListPointer[T any](list []*T) []T {
	var result []T
	for i := range list {
		result = append(result, *list[i])
	}
	return result
}
