# 协程池

## demo
```go
package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

// https://zhuanlan.zhihu.com/p/37754274

var sum int32

type TestData struct {
	I int32
	S string
}

func myFunc(data interface{}) {
	d := data.(TestData)
	atomic.AddInt32(&sum, d.I)
	fmt.Printf("run with %d\n", d.I)
	fmt.Println(d.S)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	testModel := 1

	runTimes := 100

	var wg sync.WaitGroup

	fmt.Println("start", time.Now().UnixNano())
	if testModel == 1 { // Use the common pool.

		defer ants.Release()

		syncCalculateSum := func() {
			demoFunc()
			wg.Done()
		}
		for i := 0; i < runTimes; i++ {
			wg.Add(1)
			_ = ants.Submit(syncCalculateSum)
		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", ants.Running())
		fmt.Printf("finish all tasks.\n")

	} else { // Use the pool with a function,

		var t TestData

		// set 10 to the capacity of goroutine pool and 1 second for expired duration.
		p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
			myFunc(i)
			wg.Done()
		})
		defer p.Release()
		// Submit tasks one by one.
		for i := 0; i < runTimes; i++ {
			t.I = int32(i)
			t.S = "Test" + strconv.Itoa(i)
			wg.Add(1)
			_ = p.Invoke(t) // 调用并传参
		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", p.Running())
		fmt.Printf("finish all tasks, result is %d\n", sum)

	}
	fmt.Println("end", time.Now().UnixNano())
}
```