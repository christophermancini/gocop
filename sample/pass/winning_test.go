package pass

import (
	"testing"
)

func TestWillPass(t *testing.T) {
	if myNumber() != 11 {
		t.Error("number does not equal eleven")
	}
}
