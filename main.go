package main

import (
	"fmt"
	"io"
	"os"

	"howett.net/plist"
)

func main() {
	filePath := os.Args[1]
	itermcolors, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer itermcolors.Close()
	itermbytes, _ := io.ReadAll(itermcolors)

	colors := make(map[string]map[string]interface{})
	_, err = plist.Unmarshal(itermbytes, &colors)

	for k, ansi := range colors {
		var hex string
		var red, green, blue int

		for k, component := range ansi {
			switch component := component.(type) {
			case float64:
				real := component
				switch k {
				case "Red Component":
					red = int(real * 255)
				case "Green Component":
					green = int(real * 255)
				case "Blue Component":
					blue = int(real * 255)
				}
			}
		}

		hex = fmt.Sprintf("#%02x%02x%02x", red, green, blue)
		fmt.Printf("%s: %s\n", k, hex)
	}

}
