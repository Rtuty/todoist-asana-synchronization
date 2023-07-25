package tools

func AsRef[T any](a T) *T { return &a }
