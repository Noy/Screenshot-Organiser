package main

import (
	"io/ioutil"
	"log"
	"os"
	"math/rand"
	"time"
)

var randLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil { log.Fatal(err)}
	for _, file := range files {
		if file.Name() == ".idea" {
			continue
		}
		if file.Name() == "main" {
			continue
		}
		os.Rename(file.Name(), randStringRunes(3) +".jpg")
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randLetters[rand.Intn(len(randLetters))]
	}
	return string(b)
}
