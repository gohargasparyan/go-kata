package v1

import (
	common "github.com/gohargasparyan/go-kata/karate_chop/common"
	"testing"
)

func TestChop(t *testing.T) {
	common.TestKarateChop(t, Chop)
}
