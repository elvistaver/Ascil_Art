package main

import (
	"Ascil_Art/Filesystem"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage go run . bannerfile.txt input | cat -e")
	}
	input := strings.TrimSpace(os.Args[1])

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	files := Filesystem.AscilArt("bannerFiles")

	for _, chr := range input {
		if chr == '\n' {
			fmt.Println()
			continue
		}
		if chr < 32 || chr > 126{
			return
		}

		for i := 0; i < 8; i++ {
			var line strings.Builder

			for _, ch:= range chr {
				start := (int(chr)-32)*9 + 1
				line.WriteString(files[start+i])
			}
			fmt.Print(line.String(), "\n")
		}
	}
}
