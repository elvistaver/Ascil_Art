package Filesystem

import (
	"bufio"
	"fmt"
	"os"
)

func AscilArt(input string) []string {
	result := []string{}

		file, err := os.Open("bannerFiles/standard.txt")
		if err != nil {
			fmt.Println("file not found")
		}
		readingfile := bufio.NewScanner(file)
		for readingfile.Scan() {
			line := readingfile.Text()
			result = append(result, line)
		}
		defer file.Close()
	return result	
}
