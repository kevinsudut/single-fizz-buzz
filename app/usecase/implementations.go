package usecase

import (
	"context"
	"fmt"
	"strings"
	"sync"

	singlefizzbuzz "github.com/kevinsudut/single-fizz-buzz/app/single-fizz-buzz"
)

func (u usecase) UseCaseSingleFizzBuzzWithRange(ctx context.Context, from int64, to int64) (resp string, err error) {
	if from > to || to-from > 100 {
		return resp, fmt.Errorf("invalid range")
	}

	var (
		index        = int64(0)
		goroutineSem = make(chan struct{}, 1000) // maximum goroutine
		results      = make([]string, to-from+1)
	)

	var wg sync.WaitGroup

	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(num int64, idx int64) {
			goroutineSem <- struct{}{}
			defer func() {
				<-goroutineSem
				wg.Done()
			}()

			results[idx] = singlefizzbuzz.SingleFizzBuzz(num)
		}(i, index)
		index++
	}

	wg.Wait()

	return strings.Join(results, " "), nil
}
