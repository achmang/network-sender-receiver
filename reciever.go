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

	for {
		reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter text: ")
		frame, _ := reader.ReadString('\n')
		frame = strings.Replace(frame, "\n", "", -1)

		println(frame)

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

		decodedMsg.WriteString(getFrameContent(frame))

		if isEndFrame(frame) {
			println(decodedMsg.String())
			break
		}
	}

}

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

func errorCheckSum(frame string) bool {
	if checkSum(frame[1:len(frame)-3]) == frame[len(frame)-3:len(frame)-1] {
		return true
	}
	return false
}

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
		println("4th")
		return false
	}
	return true
}

// func errorCheckFrame(frame string, mtu int) bool {

// }

func isEndFrame(frame string) bool {
	if frame[1:2] == "F" {
		return true
	}
	return false
}

func getFrameContent(frame string) string {
	return frame[6 : len(frame)-4]
}

func checkSum(frame string) string {
	total := 0

	for _, c := range frame {
		total += int(c)
	}

	hex := fmt.Sprintf("%02x", total)
	return hex[len(hex)-2:]
}
