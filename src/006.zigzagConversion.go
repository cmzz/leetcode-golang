package main

import "strings"

func convert(s string, numRows int) string {
	j := 1
	k := 0
	i := 0
	f := 0

	for i < len(s) {
		if f == 0 && j <= numRows {

			i++
		}

		if j == numRows {
			f = 1
		}
	}
}

func main() {

}