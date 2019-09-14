package main

import (
	"fmt"
	"os"
	"strconv"
)

var totalCarbon, _ = strconv.ParseInt(os.Args[1], 10, 32)
var seqMin, seqMax int64 = 4, 10

func main() {
	minSqkm := totalCarbon / seqMax
	maxSqkm := totalCarbon / seqMin
	fmt.Printf(
		"Require between %d km2 and %d km2 to sequester %d tonnes of carbon\n",
		minSqkm, maxSqkm, totalCarbon,
	)
}
