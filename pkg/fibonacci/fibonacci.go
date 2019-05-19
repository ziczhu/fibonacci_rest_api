package fibonacci

import (
	"math/big"
	"sync"
)

var (
	zeroSequence []*big.Int = nil
	oneSequence             = []*big.Int{big.NewInt(0)}
	twoSequence             = []*big.Int{big.NewInt(0), big.NewInt(1)}
)

type Fibonacci struct {
	cache        []*big.Int
	maxCacheSize int
	sync.Mutex
}

// New initializes a fillFibonacciAfterN with a cache, The big.Int uses much
// more memory than int64, but we have to use big.Int because even we
// use the uint on 64bit machine, the 94th number will overflow.
// A 10000 length fillFibonacciAfterN sequence allocates more than 50MiB memory
// on heap. Which means we should carefully setup the size and limit the input n.
// Params
// initCacheSize: The initial size of the cache, which generates the
// sequence when creates.
// maxCacheSize: The maxSize of the cache, even the maxSize of an array
// could be integer and the max value of big.Int also depends on
// the underline array, but your memory is the real limit here.
func New(initCacheSize, maxCacheSize int) *Fibonacci {
	if maxCacheSize < initCacheSize {
		panic("maxCacheSize should not small than initCacheSize")
	}

	if initCacheSize < 0 {
		panic("initCacheSize should larger than 0")
	}

	if maxCacheSize < 0 {
		panic("maxCacheSize should larger than 0")
	}

	// fixes the initCacheSize and maxCacheSize to 2 to make sure
	// the lastTwo and lastOne of cache always exists.
	if initCacheSize < 2 {
		initCacheSize = 2
	}
	if maxCacheSize < 2 {
		maxCacheSize = 2
	}

	cache := make([]*big.Int, initCacheSize, initCacheSize)
	cache[0] = big.NewInt(0)
	cache[1] = big.NewInt(1)

	if initCacheSize > 2 {
		fillFibonacciAfterN(&cache, 2)
	}

	return &Fibonacci{
		cache:        cache,
		maxCacheSize: maxCacheSize,
	}
}

// GetSequence accepts the input n, and returns first n sequence
// Internally, it will return the exist range of cache if the
// given n is smaller than the current size, otherwise it will grow
// up to at most maxSize and update the cache.
func (fib *Fibonacci) GetSequence(n int) []*big.Int {
	// we will returns the result directly if n <= 2
	if n <= 0 {
		return zeroSequence
	} else if n == 1 {
		return oneSequence
	} else if n == 2 {
		return twoSequence
	}

	size := len(fib.cache)
	// returns the range in cache if it's smaller
	if n <= size {
		ret := fib.cache[:n]
		return ret
	}

	ret := make([]*big.Int, n, n)
	for i := 0; i < size; i++ {
		ret[i] = fib.cache[i]
	}
	fillFibonacciAfterN(&ret, size)
	fib.updateCache(&ret)

	return ret
}

// update the cache, at most maxCacheSize.
func (fib *Fibonacci) updateCache(newCache *[]*big.Int) {
	fib.Lock()
	defer fib.Unlock()

	originalSize := len(fib.cache)
	newSize := len(*newCache)

	// returns immediately if already reach the maxCacheSize or
	// updated by other request
	if originalSize >= fib.maxCacheSize || originalSize >= newSize {
		return
	}

	if newSize <= fib.maxCacheSize {
		// replace the cache if the newSize small or equal than the max cache size
		fib.cache = *newCache
		return
	}

	// otherwise only update the maxCacheSize cache
	fib.cache = (*newCache)[:fib.maxCacheSize]
}

// fillFibonacciAfterN is the internal algorithms which fill the n-after
// sequence after second, the n should not smaller than 0
func fillFibonacciAfterN(nums *[]*big.Int, n int) {
	if n < 2 {
		panic("the n should not smaller than 2")
	}

	lastTwo := new(big.Int).Set((*nums)[n-2]) // clone the last two
	lastOne := new(big.Int).Set((*nums)[n-1]) // clone the last one

	for i := n; i < len(*nums); i++ {
		lastTwo.Add(lastTwo, lastOne)
		lastOne, lastTwo = lastTwo, lastOne
		(*nums)[i] = new(big.Int).Set(lastOne)
	}
}
