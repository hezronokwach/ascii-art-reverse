package reverse

import (
	"fmt"
	"os"
	"strings"
)

//read contents line by line
func ReadFile(fileName string) ([]string, bool) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error %d\n", err)
		return nil, true
	}
	var lines []string
	if fileName == "banners/thinkertoy.txt" {
		lines = strings.Split(string(content), "\r\n")
	} else {
		lines = strings.Split(string(content), "\n")
	}
	return lines, false
}
