package mascot_test

import (
	"testing"

	"github.com/caskeychr/learngo/misc/test-app/mascot"
)

func TestMascot(t *testing.T) {

	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot: (")
	}
}
