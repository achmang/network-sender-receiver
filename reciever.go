package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	decodedMsg := strings.Builder{}
	mtu := setScriptVariables()
	reader := bufio.NewScanner(os.Stdin)
	//This way of reading input it was better and readable.
	for reader.Scan() {
		frame := reader.Text()
		
		// bool logic here works but reads weirdly, should probably flip bool values.
		// reading if error makes more sense.
		// check the frame for any issues.		
		if !errorFormat(frame) {
			log.Fatal("Frame Formatting Error")
		}
		if !errorMtu(frame, mtu) {
			log.Fatal("MTU Error")
		}
		if !errorCheckSum(frame) {
			log.Fatal("Checksum Error")
		}
		if !errorMsgLen(frame) {
			log.Fatal("Msg Length Error")
		}
// 		If the frame checks out, write the content of the frame to the stringbuilder
		decodedMsg.WriteString(getFrameContent(frame))
		
// 		On receiving a final frame, terminate the program.
		if isEndFrame(frame) {
			println(decodedMsg.String())
			break
		}
	}

}
// Handle user input when running script
func setScriptVariables() int {
	mtuInput := flag.Int("mtu", 0, "Set the maximum transferable rate, mtu will be set to 30 if not specified.")
	flag.Parse()

	if *mtuInput == 0 {
		log.Fatal("MTU Not specified. Run script with an MTU parameter.")
	}

	if *mtuInput < 10 {
		log.Fatal("MTU Value must be at least 10.")
	}

	return *mtuInput
}
// Make sure the checksum matches. Issues with checksum can indicate errors in the frame.
func errorCheckSum(frame string) bool {
	if checkSum(frame[1:len(frame)-3]) == frame[len(frame)-3:len(frame)-1] {
		return true
	}
	return false
}
// Make sure the MTU constrainst is met.
func errorMtu(frame string, mtu int) bool {

	if isEndFrame(frame) {
		if len(frame) < mtu+1 {
			return true
		}
	} else {

		if mtu > 108 && len(frame) == 109 {
			return true
		} else if mtu == len(frame) {
			return true
		}
	}

	return false
}
// Make sure message length is all good
func errorMsgLen(frame string) bool {
	frameMsgLen, err := strconv.Atoi(frame[3:5])

	if err != nil {
		log.Fatal(err)
	}
	if len(getFrameContent(frame)) == frameMsgLen {
		return true
	}
	return false
}
// Frame may be malformed.
// Regex can work here but I think this is quicker?
func errorFormat(frame string) bool {
	if frame[:1] != "[" && frame[len(frame)-1:] != "]" {
		return false
	}

	if frame[1:3] != "D~" && frame[1:3] != "F~" {
		return false
	}

	_, err := strconv.Atoi(frame[3:5])
	if err != nil {
		return false
	}

	if frame[len(frame)-4:len(frame)-3] != "~" {
		return false
	}
	return true
}
// Check whether frame is a delimeter or endframe
func isEndFrame(frame string) bool {
	if frame[1:2] == "F" {
		return true
	}
	return false
}
// returns the core content of the frame.
func getFrameContent(frame string) string {
	return frame[6 : len(frame)-4]
}
// Calculate a checksum of given frame
func checkSum(frame string) string {
	total := 0

	for _, c := range frame {
		total += int(c)
	}

	hex := fmt.Sprintf("%02x", total)
	return hex[len(hex)-2:]
}
