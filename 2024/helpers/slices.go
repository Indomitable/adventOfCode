package helpers

func MapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	res := make([]R, len(slice))
	for i, item := range slice {
		res[i] = mapper(item)
	}
	return res
}

func FilterSlice[T any](slice []T, filter func(T) bool) []T {
	res := make([]T, 0)
	for _, item := range slice {
		if filter(item) {
			res = append(res, item)
		}
	}
	return res
}

func Repeat[T any](value T, count int) []T {
	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = value
	}
	return res
}

func Copy[T any](slice []T) []T {
	return append([]T{}, slice...)
}
