package main

import (
	"math/rand"
	"time"
	"io/ioutil"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")


	
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

func writeshowme(str string){
	println(str)
	d1 := []byte(str)
	println(d1)
    err := ioutil.WriteFile("/home/ec2-user/environment/showme.txt", d1, 0644)
    check(err)
}


func main() {

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	randstr := RandomString(20)


	for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				str := t.Format("2006-01-02 15:04:05") + " : " + randstr
				writeshowme(str)
			}
	}

}
