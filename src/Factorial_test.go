package testpack1

import (
	"fmt"
	"math/rand"
	"testing"
)

func facto(term int) int {
	factorial := 1
	for i := term; i > 1; i-- {

		factorial *= i

	}
	return factorial

}

func TestFacto(t *testing.T) {
	fi := facto(5)
	if fi != 120 {
		t.Error("expected ", 120, " found ", fi)
	}
}

func Examplefacto() {
	fmt.Println(facto(6))
	//Output:
	//720
}

func BenchmarkFacto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		facto(rand.Intn(3))
	}
}
