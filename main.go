package main

import (
	"Ascil_Art/Filesystem"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("invalid commandline argument: usage; go run . input text  bannerfile-type")
		return
	}
	input := strings.TrimSpace(os.Args[1])
	FileName := (os.Args[2])

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	files := Filesystem.ReadBanner(FileName)
		for i := 0; i < 8; i++ {
			line := ""
			for _, ch := range input {
				if ch < 32 || ch > 126 {
					log.Fatal(" \n not a valid character")
					return
				}
				start := (int(ch)-32)*9 + 1
				line += files[start+i]
			}
			fmt.Print(line, "\n")
		}
	
}
