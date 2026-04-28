package Filesystem

import (
	"bufio"
	"fmt"
	"os"
)

func AscilArt(input string) []string {
	final := []string{}

	data, err := os.ReadDir(input)

	if err != nil {
		fmt.Println("file not in directory")
	}

	for _, files := range data {
		fs := files.Name()
		if files.IsDir() {
			continue
		}
		tex, err := os.Open("bannerFiles/" + fs)
		if err != nil {
			fmt.Println("file not found")
		}
		text := bufio.NewScanner(tex)
		for text.Scan() {
			line := text.Text()
			final = append(final, line)
		}
		defer tex.Close()
	}
	return final	
}
