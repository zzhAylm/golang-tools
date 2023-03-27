package morestrings

import (
	"fmt"
	"testing"
)

func FuzzHello(f *testing.F) {
	hello := Hello("zzh")

	fmt.Println(hello)
}
