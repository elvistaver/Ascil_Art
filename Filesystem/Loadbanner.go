package Filesystem

import (
	"bufio"
	"log"
	"os"
)

func ReadBanner(bannerName string) []string {
	result := []string{}

	file, err := os.Open("bannerFiles/" + bannerName + ".txt")
	if err != nil {
		log.Fatal("\nfile not found")
	}
	defer file.Close()

	readingfile := bufio.NewScanner(file)
	for readingfile.Scan() {
		line := readingfile.Text()
		result = append(result, line)
	}
	// if err == readingfile.Err(){
	// 	log.Fatal("\n could not read File")
	// }
	return result
}
