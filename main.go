package main

import (
	"Ascil_Art/Filesystem"
	"fmt"
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
	inputfile:= strings.Split(input, "\\n")

	files := Filesystem.ReadBanner(FileName)

	for _, chr:= range inputfile{
		if chr == ""{
			fmt.Println()
			continue	
		}
			for i := 0; i < 8; i++ {
				line := ""
				for _, ch := range chr {
					if ch < 32 || ch > 126 {
						fmt.Printf(" \n not a valid character%v",string(ch))
						return
					}
					start := (int(ch)-32)*9 + 1
					line += files[start+i]
				}
				fmt.Print(line, "\n")
			}
	}

}
