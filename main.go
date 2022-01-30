package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	starttime := time.Now()
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic("there was an error opening the file, we can't proceed")
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	var inputlines []string
	for reader.Scan() {

		inputlines = append(inputlines, reader.Text())

	}
	numbers := strings.Split(inputlines[0], " ")
	var LR []string
	for i := 1; i < len(inputlines); i++ {
		LR = strings.Split(inputlines[i], " ")

		L, err_L := strconv.Atoi(LR[0])
		R, err_R := strconv.Atoi(LR[1])
		if err_L != nil || err_R != nil {
			log.Printf("Some of the value are not integers")

		}
		sum := 0
		for j := L - 1; j < R; j++ {
			number, err := strconv.Atoi(numbers[j])
			if err != nil {
				log.Printf("Some of the value are not integers")
			}
			sum += number

		}
		fmt.Printf("[L,R] is : [%s,%s], Sum is : %d\n", LR[0], LR[1], sum)

	}
	excutiontime := time.Since(starttime)
	fmt.Printf("Excution time : %s", excutiontime)

}
