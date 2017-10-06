package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

// check if filepath is only argument passed
// if so print help
// if not, pass arguments along to proper functions

func sha256Sum(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return hash.Sum(nil)

}

func sha1Sum(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return hash.Sum(nil)

}

func md5Sum(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return hash.Sum(nil)
}

func printHelp() {
	help := "usage: chksm [algorithm] [file]\n" +
		"algorithms available: md5, sha1, sha256\n"
	fmt.Printf("%s", help)
}

func main() {

	if len(os.Args) == 3 {
		hashAlgo := os.Args[1]
		filename := os.Args[2]
		switch hashAlgo {
		case "md5":
			newSum := md5Sum(filename)
			fmt.Printf("%s: %x", filename, newSum)
		case "sha1":
			newSum := sha1Sum(filename)
			fmt.Printf("%s: %x", filename, newSum)
		case "sha256":
			newSum := sha256Sum(filename)
			fmt.Printf("%s: %x", filename, newSum)
		default:
			printHelp()
			// Print help because the supplied algo isn't available
		}
	} else {
		printHelp()
		// Print help because too few or too many arguments were given
	}

}
