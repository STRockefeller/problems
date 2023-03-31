package random

import (
	"testing"

	linq "github.com/STRockefeller/go-linq"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAmountOf(t *testing.T) {
	assert := assert.New(t)
	randomSlice := randomStringSlice(fake.Uint8())
	amount := fake.IntRange(0, len(randomSlice))
	actual := AmountOf(randomSlice, amount)
	assert.Len(actual, amount)
	linq.Linq[string](actual).ForEach(func(s string) {
		assert.True(linq.Linq[string](randomSlice).Contains(s))
	})
}

func randomStringSlice(amount uint8) []string {
	res := make([]string, amount)
	for i := 0; i < int(amount); i++ {
		res[i] = fake.CarType()
	}
	return res
}

func TestOneOf(t *testing.T) {
	assert := assert.New(t)
	randomSlice := randomStringSlice(fake.Uint8())
	actual := OneOf(randomSlice)
	assert.True(linq.Linq[string](randomSlice).Contains(actual))
}
