package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	hash := md5.New()
	body, err := ioutil.ReadFile("md5.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal err : %v ", err)
		os.Exit(1)
	}
	hash.Write(body)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()
	fmt.Printf("hash Value : %s, size : %d\n", string(hashValue), hashSize)
	for i := 0; i < hashSize; i += 4 {
		var val uint32
		val = uint32(hashValue[i])<<24 +
			uint32(hashValue[i+1])<<16 +
			uint32(hashValue[i+2])<<8 +
			uint32(hashValue[i+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
