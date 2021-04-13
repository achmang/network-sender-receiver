package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Get the user script variables.
	mtuVal, fileParameter := scriptVariables()
	// Check if user wants to read file as input.
	if fileParameter != "" {
		encodeMessage(fileParameter, mtuVal)
	}

	for {
		// This will print weird on windows, optimised for linux.
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		encodeMessage(text, mtuVal)
	}
}

//Handle users script variables
func scriptVariables() (int, string) {

	fileInput := flag.String("file", "", "Type the name of the file that you want to read.")
	mtuInput := flag.Int("mtu", 0, "Set the maximum transferable rate, mtu will be set to 30 if not specified.")
	flag.Parse()

	f, err := ioutil.ReadFile(*fileInput)

	if *fileInput != "" {
		if err != nil {
			log.Fatal(err)
		}
	}

	if *mtuInput == 0 {
		log.Fatal("MTU Not specified. Run script with an MTU parameter.")
	}

	if *mtuInput < 10 {
		log.Fatal("MTU Value must be at least 10.")
	}

	return *mtuInput, string(f)
}

func encodeMessage(msg string, mtu int) {

	// if the frame is of length 10 then it cannot have message content
	if len(msg) > 0 && mtu == 10 {
		log.Fatal("Frame length exceeded.")
	}

	emptyFrameSize := 10
	msgSize := len(msg)
	stringEndIndex := mtu - emptyFrameSize
	stringStartIndex := 0

	// content inside a frame cannot be greater than 99, so set this limit
	if stringEndIndex > 99 {
		stringEndIndex = 99
	}

	// 
	for stringEndIndex < msgSize {
		frame := frameSubstring(msg[stringStartIndex:stringEndIndex], "D")

		fmt.Println(frame)
		if stringEndIndex >= 99 && mtu >= 109 {
			stringEndIndex += 99
			stringStartIndex += 99
		} else {
			stringEndIndex += mtu - emptyFrameSize
			stringStartIndex += mtu - emptyFrameSize
		}

	}

	frame := frameSubstring(msg[stringStartIndex:msgSize], "F")
	fmt.Println(frame)

}

func frameSubstring(subMsg string, frameType string) string {

	builder := strings.Builder{}

	builder.WriteString("[")
	builder.WriteString(frameType)
	builder.WriteString("~")
	builder.WriteString(fmt.Sprintf("%02d", len(subMsg)))
	builder.WriteString("~")
	builder.WriteString(subMsg)
	builder.WriteString("~")
	builder.WriteString(checkSum(builder.String()[1:]))
	builder.WriteString("]")

	return builder.String()
}

func checkSum(frame string) string {

	total := 0

	for _, c := range frame {
		total += int(c)
	}

	hex := fmt.Sprintf("%02x", total)

	return hex[len(hex)-2:]
}
