package tern

// E stands for evaluate
func E[T interface{}](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
