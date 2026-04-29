package main

import (
	"Ascil_Art/Filesystem"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	input := strings.TrimSpace(os.Args[1])

	if len(os.Args) != 3 {
		fmt.Println("leave")
		return
	}

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	files := Filesystem.ReadBanner("bannerFiles/.txt")

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
