package main

import (
	"fmt"
	"strconv"
	"strings"
)

type productRanges struct {
	first int
	last  int
}

func day02a(rawdata []string) {
	var invalidProductIDs []string
	answer := 0
	ProductIDRanges := processDay2Data(rawdata[0])

	for _, prange := range ProductIDRanges {
		var productIDlist []int
		for productid := prange.first; productid <= prange.last; productid++ {

			productIDlist = append(productIDlist, productid)

		}
		//fmt.Println(productIDlist)
		for _, product := range productIDlist {
			ptext := strconv.Itoa(product)
			if len(ptext)%2 == 0 {
				if ptext[:len(ptext)/2] == ptext[len(ptext)/2:] {
					invalidProductIDs = append(invalidProductIDs, ptext)
				}
			}
		}
	}
	for _, product := range invalidProductIDs {
		productnum, _ := strconv.Atoi(product)
		answer = answer + productnum

	}
	fmt.Println("Day 02a Answer: ", answer)
}

func processDay2Data(rawdata string) []productRanges {
	var ResultData []productRanges
	idRanges := strings.Split(rawdata, ",")
	for _, ranges := range idRanges {
		//fmt.Println(ranges)
		rangesplit := strings.Split(ranges, "-")
		a, _ := strconv.Atoi(rangesplit[0])
		b, _ := strconv.Atoi(rangesplit[1])
		var idrange productRanges
		idrange.first = a
		idrange.last = b
		ResultData = append(ResultData, idrange)

	}
	return ResultData
}
