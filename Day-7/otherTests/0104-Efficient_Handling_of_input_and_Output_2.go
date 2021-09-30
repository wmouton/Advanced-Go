package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

func main() {
	// Bytes from a png image or whatever
	pngPayload := []byte{137, 80, 78, 71, 13, 10, 26, 10, 11, 12, 14}
	buf := make([]byte, 4)

	// read the first 4 bytes into buffer - Essentially the png header
	_, err := io.ReadFull(bytes.NewReader(pngPayload), buf)
	if err != nil {
		log.Fatal("error reading png data")
	}

	fmt.Println(buf)

	// alternative way to write to stdout
	io.WriteString(os.Stdout, string(buf))

	// alternative implementation
	lr := io.LimitReader(bytes.NewReader(pngPayload), 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	tpmFile, err := ioutil.TempFile(".", "temp_")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.Remove(tmpFile.Name()) // clean up
	defer func(tmpFile *os.File) {
		// Close can fail
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
	}(tmpFile)
 }