package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 10; i++ {
		count := i // Account for base 1
		rep := make([]string, i)
		rep[0] = strings.Repeat("1", i)

		for base := 2; base <= i; base++ {
			s := strconv.FormatInt(int64(i), base)

			rep[base-1] = s

			count += len(s)
		}

		fmt.Printf("%d: %d - %s\n", i, count, strings.Join(rep, " "))
	}
}
