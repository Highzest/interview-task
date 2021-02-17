package decrypt

import (
	"flag"
	"fmt"
)

func main() {
	textPtr := flag.String("t", "Hello, world!", "Text to decrypt.")
	keyPtr := flag.Int("k", 3, "Encryption/decryption key.")
	flag.Parse()

	fmt.Printf("textPtr: %s, keyPtr: %d", *textPtr, *keyPtr)
}
