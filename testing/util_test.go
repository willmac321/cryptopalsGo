package challenge

import (
	utils "cryptoGo/util"
	"testing"
)

func TestGetEnd(t *testing.T) {
	exp := 7
	res := utils.GetEnd(3, 4, 10)
	if res != exp {
		t.Error("Expected ", exp, "\n but got ", res)
	}
	exp1 := 6
	res1 := utils.GetEnd(3, 3, 9)
	if res1 != exp1 {
		t.Error("Expected ", exp, "\n but got ", res)
	}
}
