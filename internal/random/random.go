package random

import fake "github.com/brianvoe/gofakeit/v6"

func AmountOf[T any](items []T, amount int) []T {
	if amount <= 0 {
		return []T{}
	}
	length := len(items)
	if amount >= length {
		return items
	}
	result := make([]T, amount)
	for i := 0; i < amount; i++ {
		index := fake.IntRange(i, length-1)
		result[i] = items[index]
		items[i], items[index] = items[index], items[i] // remove selected item from slice
	}
	return result
}

func OneOf[T any](items []T) T {
	length := len(items)
	if length == 0 {
		return *new(T)
	}
	return items[fake.IntRange(0, length-1)]
}
