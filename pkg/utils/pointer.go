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

func WrapPointerInt64(value int64) *int64 {
	return &value
}

func UnwrapPointerInt64(value *int64) int64 {
	return *value
}

func WrapPointerInt32(value int32) *int32 {
	return &value
}

func UnwrapPointerInt(value *int) int {
	return *value
}

func WrapPointerInt(value int) *int {
	return &value
}

func UnwrapPointerInt32(value *int32) int32 {
	return *value
}
