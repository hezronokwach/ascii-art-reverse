package arguments

import (
	"ascii-art-reverse/reverse"
	"fmt"
	"os"

	"strings"
)

func Args() {
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
		return
	}
	arg := args[0]
	banner := "banners/"

	switch len(args) {
	case 1:
		banner += "standard.txt"
	case 2:
		banner = args[1]
	}

	if !strings.Contains(arg, "--reverse=") {
		printUsage()
		return
	}
	bannerFile, err := reverse.ReadFile(banner)
	if err {
		printUsage()
		return
	}

	sliceadBanner := reverse.SliceFile(bannerFile)
	file, err := reverse.ReadFile(arg[10:])
	if err {
		printUsage()
		return
	}

	str := ""
	for len(file[0]) > 0 {
		for i, val := range sliceadBanner {
			if reverse.CheckPattern(val, file) {
				str += string(rune(i + 32))
				file = reverse.TrimFound(len(val[0]), file)
			}
		}
	}
	fmt.Println(str)
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
}
