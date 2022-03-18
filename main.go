package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Missing file argument. Usage: %s <file> <lines>", os.Args[0])
	}

	lines := 10
	if len(os.Args) > 2 {
		l, _ := strconv.Atoi(os.Args[2])
		if l <= 0 {
			log.Fatal("invalid lines arguments")
		}
		lines = l
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer f.Close()

	// seek to end of file
	off, err := f.Seek(0, 2)
	if err != nil {
		log.Fatalf("seek error: %s", err)
	}

	const newLine byte = '\n'

	buf := make([]byte, 1)
	var foundLines int
	lastChar := true

	// search new lines backwards
	for off > 0 && foundLines < lines {
		off = off - 1
		if _, err = f.ReadAt(buf, off); err != nil {
			log.Fatalf("read error: %s", err)
		}
		// if last char is a new line, doesn't count
		if buf[0] == newLine && !lastChar {
			foundLines++
		}
		lastChar = false
	}

	// increase +1 so not to print first new line
	if off > 0 {
		off += 1
	}

	// seek to starting point
	f.Seek(off, 0)
	buf = make([]byte, 64)
	for {
		if _, err = f.Read(buf); err != nil {
			break
		}
		os.Stdout.Write(buf)
	}
}
