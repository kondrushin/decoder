package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kondrushin/decoder/decoder"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	color.Cyan("----------------------------------------------------------------------------------")
	color.Cyan("The application calculates the number of ways your encoded message can be decoded.")
	color.Cyan("----------------------------------------------------------------------------------")

	for {
		color.Cyan("Write your message here -> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if len(text) > 100 {
			color.Red("Sorry, your message is too long")
		} else {
			var re = regexp.MustCompile(`^[0-9]+$`)
			if re.MatchString(text) {
				number := decoder.Decode(text)
				color.Green("The numbe of ways to decode your message is -> %d\n", number)
			} else {
				color.Red("String is not valid. It should contain only digits.")
			}
		}
	}
}
