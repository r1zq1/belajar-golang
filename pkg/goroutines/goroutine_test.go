package goroutines

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func RunHello() {
	fmt.Println("Hello from Goroutines")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutines(t *testing.T) {
	fmt.Println("Entering test function")
	go RunHello()
	fmt.Println("Leaving test function")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	start := time.Now()
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	totalTime := time.Since(start)
	fmt.Println(totalTime)

	time.Sleep(10 * time.Second)
}
