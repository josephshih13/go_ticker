package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func main() {

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	randstr := RandomString(20)

	// 	go func() {
	// 		for {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case t := <-ticker.C:
	// 				fmt.Println("Tick at", t, randstr)
	// 			}
	// 		}
	// 	}()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			// 			fmt.Println("Tick at", t, randstr)
			fmt.Println(t, ": ", randstr)
		}
	}

	time.Sleep(16 * time.Second)
	// ticker.Stop()
	// done <- true
	// fmt.Println("Ticker stopped")
}
