package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s [file part 1] [file part 2] [output file]\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open part 1 file - %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	defer file.Close()
	data1, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	file, err = os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open part 1 file - %s: %s\n", os.Args[2], err)
		os.Exit(1)
	}
	defer file.Close()
	data2, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %s: %s\n", os.Args[2], err)
		os.Exit(1)
	}

	if len(data1) != len(data2) {
		fmt.Fprintf(os.Stderr, "Part 1 and part 2 have different sizes (%d != %d)", len(data1), len(data2))
		os.Exit(1)
	}

	out := make([]byte, len(data1))

	for i := range out {
		out[i] = data1[i] ^ data2[i]
	}

	file, err = os.Create(os.Args[3])
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file %s: %s", os.Args[3], err)
	}

	file.Write(out)

	fmt.Println("Done!")
}
