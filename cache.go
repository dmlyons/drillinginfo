//
// cache.go
//
// A basic cache written in Go.
// This is an exercise where we will implement a simple key value store written in Go.
// We will use a simple readline interface and two commands: PUT and GET.
//
// Requirements:
//
// 1. PUT key value     Set a value in the cache.
// 2. GET key           Get a value stored in the cache.
// 3. EXIT/QUIT         Exits the interactive prompt (can also be done with Ctrl-d thanks to the readline pkg).
// 4. Use only packages from the stdlib (except for the readline package already imported below).
//
package main

import (
	"io"
	"log"

	"github.com/chzyer/readline"
)

func main() {
	prompt, err := readline.New("> ")
	if err != nil {
		log.Fatal(err)
	}
	defer prompt.Close()

	for {
		line, err := prompt.Readline()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_ = line // TODO: implement key/value store!
	}
}
