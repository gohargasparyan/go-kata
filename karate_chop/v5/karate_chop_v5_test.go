package v1

import (
	"github.com/gohargasparyan/go-kata/karate_chop/common"
	"testing"
)

func TestChop(t *testing.T) {
	common.TestKarateChop(t, Chop)
}

func Benchmark(b *testing.B) {
	common.Benchmark(b, Chop)
}