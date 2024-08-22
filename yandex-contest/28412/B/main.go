package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type guest struct {
	Number   int
	Side     string
	Position string
}

func main() {
	var (
		numRows   int
		seats     [][]int
		numGuests int
		guests    []guest
		tmp       string

		scanner = bufio.NewScanner(os.Stdin)
	)

	/*
		    FREE = 1
			BOOKED = 2
			BYPASS = 3
	*/

	fmt.Scan(&numRows)

	seats = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		fmt.Scan(&tmp)

		seats[i] = make([]int, 7)

		for j, e := range tmp {
			if e == '.' {
				seats[i][j] = 1
			} else if e == '#' {
				seats[i][j] = 2
			} else if e == '_' {
				seats[i][j] = 3
			}
		}
	}

	fmt.Scan(&numGuests)

	guests = make([]guest, numGuests)
	for i := 0; i < numGuests; i++ {
		guests[i] = getGuest(scanner)

		// search for suitable side
		for _, e := range seats {
			var sideSeats []int
			if guests[i].Side == "left" {
				sideSeats = e[0:3]
			} else {
				sideSeats = e[4:7]
			}

			if guests[i].Position == "aisle" {

			}
		}
		// search for suitable number according to side
		// search for suitable position
	}
}

func getGuest(scanner *bufio.Scanner) guest {
	scanner.Scan()
	tmp := scanner.Text()
	tmpList := strings.Split(tmp, " ")

	guestNum, _ := strconv.Atoi(tmpList[0])

	return guest{guestNum, tmpList[1], tmpList[2]}
}
