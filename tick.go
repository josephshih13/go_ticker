package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type SafeQueue struct {
	mu   sync.Mutex
	data []string
}

func (q *SafeQueue) Enqueue(s string) {
	q.mu.Lock()
	q.data = append(q.data, s)
	q.mu.Unlock()
}

func (q *SafeQueue) Getstring() string {
	q.mu.Lock()
	ret := ""

	start := 0
	if len(q.data) > 5 {
		start = len(q.data) - 5
	}

	for cnt, idx := 0, start; idx < len(q.data) && cnt < 5; idx, cnt = idx+1, cnt+1 {
		ret += q.data[idx] + "\n"
	}

	// if len(q.data) == 0 {
	// 	q.mu.Unlock()
	// 	return ""
	// }
	// v := q.data[0]
	q.data = q.data[start:]
	q.mu.Unlock()
	return ret
}

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
	queue := SafeQueue{data: []string{}}

	randstr := RandomString(20)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				str := t.Format("2006-01-02 15:04:05") + ": " + randstr
				queue.Enqueue(str)
			}
		}
	}()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, queue.Getstring())
	})

	// Start server
	fmt.Println("Start Server from port 9936")
	e.Logger.Fatal(e.Start(":9936"))
}
