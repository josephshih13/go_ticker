package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var outputstr string

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func getpong() string {
	resp, err := http.Get("http://ping-pong-svc:6789/")
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	// print(body)
	return string(body)
}

func main() {

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	randstr := RandomString(20)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				str := t.Format("2006-01-02 15:04:05") + " : " + randstr + "\n"
				outputstr = str
			}
		}
	}()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, outputstr+getpong())
	})

	// Start server
	fmt.Println("Start Server from port 9936")
	e.Logger.Fatal(e.Start(":9936"))

}
