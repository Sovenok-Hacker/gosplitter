package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [file to split]\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	rd := make([]byte, len(content))
	rand.Read(rd)

	xrd := make([]byte, len(content))
	for i := range xrd {
		xrd[i] = rd[i] ^ content[i]
	}

	f1, err1 := os.Create(fmt.Sprintf("%s.gosplitter.1", os.Args[1]))
	f2, err2 := os.Create(fmt.Sprintf("%s.gosplitter.2", os.Args[1]))

	defer f1.Close()
	defer f2.Close()

	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Error opening output file: %s\n", err1)
		os.Exit(1)
	}

	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Error opening output file: %s\n", err2)
		os.Exit(1)
	}

	f1.Write(rd)
	f2.Write(xrd)

	fmt.Println("Done!")
}
