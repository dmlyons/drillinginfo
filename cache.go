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
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/chzyer/readline"
)

var errCacheMiss = errors.New("Key not found in cache")

type cache struct {
	// readline is only ever going to give us strings, so...
	store map[string]string

	// probably not necessary for this test, but in any real application it
	// would likely be the first thing added
	lock sync.RWMutex
}

func (c *cache) Put(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.store[key] = value
}

func (c *cache) Get(key string) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if value, ok := c.store[key]; ok {
		return value, nil
	}
	return "", errCacheMiss
}

func main() {
	c := &cache{
		store: make(map[string]string),
	}

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

		// clean things up, remove any whitespace on either side
		input := strings.TrimSpace(line)

		// figure out the command
		// time to quit?
		if exitCmd := strings.ToLower(input); exitCmd == "quit" || exitCmd == "exit" {
			os.Exit(0)
		}

		// something else, pull off the command, key, and possibly the value,
		// making the assumption here that keys do not contain spaces so as to
		// avoid parsing through quoted strings
		parts := strings.SplitN(input, " ", 3)
		// do the command
		switch strings.ToLower(parts[0]) {
		case "put":
			// well formed put has 3 parts
			if len(parts) == 3 {
				c.Put(parts[1], parts[2])
				fmt.Printf("Success: %s=\"%s\"\n", parts[1], parts[2])
			} else {
				fmt.Println("Invalid PUT: Syntax is PUT <KEY> <VALUE>")
				fmt.Println("Keys must not contain spaces")
			}
		case "get":
			// well formed get has 2 parts
			if len(parts) == 2 {
				val, err := c.Get(parts[1])
				if err != nil {
					if err == errCacheMiss {
						fmt.Printf("Key %s not found in cache\n", parts[1])
						continue
					}
					// this should be impossible with this implementation
					log.Fatal(err)
				}
				fmt.Println(val)
			} else {
				fmt.Println("Invalid GET: Syntax is GET <KEY>")
				fmt.Println("Keys must not contain spaces")
			}

		default:
			// Something is wrong here
			fmt.Printf("Invalid command: \"%s\"\n", input)
		}
	}
}
