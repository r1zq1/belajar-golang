package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("before send data")
		channel <- "Ini data yang dikirim"
		fmt.Println("after send data")
	}()
	data := <-channel
	fmt.Println(data)
	close(channel)
}

func GiveMeData(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Data 123456"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	go GiveMeData(channel)
	go GiveMeData(channel)
	data := <-channel
	fmt.Println(data)
	// data2 := <-channel
	// fmt.Println(data2)
	close(channel)
}
