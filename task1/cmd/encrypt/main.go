package main

import (
	"flag"
	"fmt"
	"log"

	cipher "github.com/Highzest/interview-task/task1"
)

func main() {
	textPtr := flag.String("t", "Hello, world!", "Text to encrypt. Spaces are not allowed.")
	keyPtr := flag.Int("k", 3, "Encryption/decryption key.")
	flag.Parse()

	if *keyPtr < 0 || *keyPtr > 26 {
		log.Fatalf("key should be between 0 and 26. Not %d", *keyPtr)
	}

	fmt.Println(cipher.Encrypt(*textPtr, *keyPtr))
}
