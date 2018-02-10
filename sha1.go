package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var hashes = make(map[string]string, 8)

var start = time.Now().UTC()

func main() {
	// Initialize
	hashes["D35DB127DB631E6E27C6B75E8D376B04F64FAF83"] = "root"
	hashes["C2D24DCA38E9E862098B85BF0AB35CAA52803797"] = "pma"
	hashes["81101DED975D54BD76A3C8EAD293597AE9BB143F"] = "Alice"
	hashes["1FDB0D828172183735F1ED9E45E6AF3CE04DE9D1"] = "Bob"
	hashes["2470C0C06DEE42FD1618BB99005ADCA2EC9D1E19"] = "Carol"
	hashes["6063C78456BB048BAF36BE1104D12D547834DFEA"] = "David"
	hashes["CE6B124C2F3D90CA4DEE06D35B81FF3A9A22EE32"] = "Evan"
	hashes["D5D82EF0B0A344A314E3B3E8D1E78A4697B97A53"] = "Flynn"

	set := "abcdefghijklmnopqrstuvwxyz"

	start = time.Now().UTC()

	// Trying passwords with 4 to 10 characters
	for i := 4; i < 11; i++ {
		go bruteForce(i, set)
	}
	go dictionaryAttack()

	// Prevent exit
	fmt.Scanln()
	fmt.Println("done")
}

// Return hashed input: sha1(sha1(input))
func hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	hash1 := h.Sum(nil)

	h = sha1.New()
	h.Write(hash1[:])
	hash2 := hex.EncodeToString(h.Sum(nil))

	output := strings.ToUpper(hash2)
	return output
}

// Return whether "input" is a password hash
func hasHash(input string) bool {
	if _, ok := hashes[input]; ok {
		return true
	}
	return false
}

func bruteForce(length int, set string) {
	np := nextPassword(length, set)
	for {
		pwd := np()
		if len(pwd) == 0 {
			break
		}
		if hasHash(hash(pwd)) {
			fmt.Println(pwd, hash(pwd), time.Since(start))
		}
	}
}

// This function return a function
// To generate all possible passwords from given set
func nextPassword(n int, c string) func() string {
	r := []rune(c)
	p := make([]rune, n)
	x := make([]int, len(p))
	return func() string {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = r[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(r) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return string(p)
	}
}

// Dictionary Attack
func dictionaryAttack() {
	lines, err := readLines("dictionary.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, word := range lines {
		if hasHash(hash(word)) {
			fmt.Println(word, hash(word), time.Since(start))
		}
	}
}

// Reads a whole file into memory
// Returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
