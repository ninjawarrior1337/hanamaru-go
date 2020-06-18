package util

import "math/rand"

func IntnNoDup(n int, prevPick *int) int {
	if r := rand.Intn(n); r == *prevPick {
		return IntnNoDup(n, prevPick)
	} else {
		return r
	}
}
