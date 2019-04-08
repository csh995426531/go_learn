package share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("counter:", counter)
}

func TestCounterMutex(t *testing.T) {

	var mt sync.Mutex
	counter := 0

	for i := 0; i < 5000; i++ {

		go func() {
			defer func() {
				mt.Unlock()
			}()
			mt.Lock()
			counter++
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("counter:", counter)
}

func TestCounterWaitGroup(t *testing.T) {

	var wg sync.WaitGroup
	var mt sync.Mutex
	counter := 0

	for i := 0; i < 500000; i++ {
		mt.Lock()
		wg.Add(1)
		go func() {
			defer func() {
				mt.Unlock()
				wg.Done()
			}()

			counter++
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
