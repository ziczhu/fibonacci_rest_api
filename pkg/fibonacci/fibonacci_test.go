package fibonacci

import (
	"sync"
	"testing"

	"math/big"

	"github.com/stretchr/testify/assert"
)

func TestGetFibonacci(t *testing.T) {
	t.Run("It should return appropriate result when n equal to 0, 1, 2", func(t *testing.T) {
		fib := New(10, 10)

		assert.Equal(t, fib.GetSequence(0), zeroSequence)
		assert.Equal(t, fib.GetSequence(1), oneSequence)
		assert.Equal(t, fib.GetSequence(2), twoSequence)
	})

	// we can only write n = 92, because n = 93 which is larger than math.MaxInt64
	t.Run("It should return appropriate result when n equal to number larger than 2", func(t *testing.T) {
		fib := New(100, 100)
		// we will test some values here
		data := []struct {
			n    int
			want *big.Int
		}{
			{0, big.NewInt(0)}, {2, big.NewInt(1)}, {3, big.NewInt(2)},
			{4, big.NewInt(3)}, {5, big.NewInt(5)}, {10, big.NewInt(55)},
			{42, big.NewInt(267914296)}, {80, big.NewInt(23416728348467685)},
			{92, big.NewInt(7540113804746346429)},
		}

		ret := fib.GetSequence(100)

		for _, d := range data {
			assert.Equal(t, ret[d.n], d.want)
		}
	})

	t.Run("It should enlarge cache result when it get larger number", func(t *testing.T) {
		initCacheSize, maxCacheSize := 10, 100
		fib := New(initCacheSize, maxCacheSize)

		assert.Equal(t, len(fib.cache), initCacheSize)
		fib.GetSequence(8)
		// get n small than current cache won't enlarge cache
		assert.Equal(t, len(fib.cache), initCacheSize)
		fib.GetSequence(11)
		// get n larger than current cache size will enlarge cache
		assert.Equal(t, len(fib.cache), 11)
		fib.GetSequence(110)
		// cache will grow up to at most the maxCacheSize
		assert.Equal(t, len(fib.cache), maxCacheSize)
	})

	t.Run("It should panic when the input is illegal", func(t *testing.T) {
		assert.Panics(t, func() {
			New(3, 1) // init should not larger than max
		})

		assert.Panics(t, func() {
			New(-3, -1) // init & max both should not smaller than 0
		})
	})

	t.Run("It should ok when the concurrency update cache", func(t *testing.T) {
		fib := New(10, 200)

		data := []struct {
			n    int
			want *big.Int
		}{
			{0, big.NewInt(0)}, {2, big.NewInt(1)}, {3, big.NewInt(2)},
			{4, big.NewInt(3)}, {5, big.NewInt(5)}, {10, big.NewInt(55)},
			{42, big.NewInt(267914296)}, {80, big.NewInt(23416728348467685)},
			{92, big.NewInt(7540113804746346429)},
		}

		var wg sync.WaitGroup

		runs := []int{100, 200, 150, 120, 180, 300}

		for _, n := range runs {
			wg.Add(1)
			go func(num int) {
				defer wg.Done()
				ret := fib.GetSequence(num)
				for _, d := range data {
					assert.Equal(t, ret[d.n], d.want)
				}
			}(n)
		}

		wg.Wait()

		// the cache will be the largest number less than maxCacheSize
		assert.Equal(t, len(fib.cache), 200)
	})
}
