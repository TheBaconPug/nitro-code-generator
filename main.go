package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unsafe"
)

func main() {
	f, err := os.Create("codes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("How many codes? ")
	var amount int
	fmt.Scanln(&amount)
	stime := time.Now()
	for i := 1; i < amount+1; i++ {
		//fmt.Println("discord.gift/" + randomString(16))
		f.WriteString("discord.gift/" + randomString(16) + "\n")
	}
	etime := time.Since(stime)
	fmt.Println("That took:")
	fmt.Println(etime.Round(time.Duration(stime.Second())))
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func randomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
