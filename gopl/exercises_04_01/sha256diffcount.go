// Write a function that counts the number of bits that are different in two
// SHA256 hashes.

package main

import (
	"crypto/sha256"
	"fmt"
)

func sha256diffcount(c1, c2 *[sha256.Size]byte) int {
	var count int
	for i := range c1 {
		diff := c1[i] ^ c2[i]
		for j := 0; j < 8; j++ {
			count += int(diff & 1)
			diff >>= 1
		}
	}

	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(sha256diffcount(&c1, &c2))

	c1 = sha256.Sum256([]byte("au"))
	c2 = sha256.Sum256([]byte("au"))
	fmt.Println(sha256diffcount(&c1, &c2))
}
